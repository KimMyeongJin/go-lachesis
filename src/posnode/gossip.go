package posnode

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Fantom-foundation/go-lachesis/src/hash"
	"github.com/Fantom-foundation/go-lachesis/src/inter"
	"github.com/Fantom-foundation/go-lachesis/src/posnode/api"
)

const (
	gossipIdle = time.Second * 5
)

// gossip is a pool of gossiping processes.
type gossip struct {
	tickets chan struct{}

	sync.Mutex
}

func (g *gossip) freeTicket() {
	g.Lock()
	defer g.Unlock()

	if g.tickets != nil {
		g.tickets <- struct{}{}
	}
}

// StartGossip starts gossiping.
func (n *Node) StartGossip(threads int) {
	n.gossip.Lock()
	defer n.gossip.Unlock()

	if n.gossip.tickets != nil {
		return
	}

	n.initPeers()

	n.gossip.tickets = make(chan struct{}, threads)
	for i := 0; i < threads; i++ {
		n.gossip.tickets <- struct{}{}
	}

	go n.gossiping(n.gossip.tickets)

	n.log.Info("gossip started")
}

// StopGossip stops gossiping.
func (n *Node) StopGossip() {
	n.gossip.Lock()
	defer n.gossip.Unlock()

	if n.gossip.tickets == nil {
		return
	}

	close(n.gossip.tickets)
	n.gossip.tickets = nil

	n.log.Info("gossip stopped")
}

// gossiping is a infinity gossip process.
func (n *Node) gossiping(tickets chan struct{}) {
	for range tickets {
		go func() {
			defer n.gossip.freeTicket()
			peer := n.NextForGossip()
			if peer != nil {
				defer n.FreePeer(peer)
				n.syncWithPeer(peer)
			} else {
				n.log.Warn("no candidate for gossip")
			}
			time.Sleep(gossipIdle)
		}()
	}

}

func (n *Node) syncWithPeer(peer *Peer) {
	client, free, fail, err := n.ConnectTo(peer)
	if err != nil {
		return
	}
	defer free()

	unknowns, err := n.compareKnownEvents(client, peer)
	if err != nil {
		fail(err)
		return
	}
	if unknowns == nil {
		return
	}

	peers2discovery := make(map[hash.Peer]struct{})
	parents := hash.Events{}

	toDownload := n.lockFreeHeights(unknowns)
	defer n.unlockFreeHeights(toDownload)

	for creator, interval := range toDownload {
		req := &api.EventRequest{
			PeerID: creator.Hex(),
		}
		for i := interval.from; i <= interval.to; i++ {
			req.Index = i

			event, err := n.downloadEvent(client, peer, req)
			if err != nil {
				fail(err)
				return
			}
			if event == nil {
				return
			}

			peers2discovery[creator] = struct{}{}
			parents.Add(event.Parents.Slice()...)
		}
	}
	n.ConnectOK(peer)

	n.checkParents(client, peer, parents)

	// check peers from events
	for p := range peers2discovery {
		n.CheckPeerIsKnown(peer.ID, peer.Host, &p)
	}

	// Clean outdated data about peers.
	n.trimHosts(n.conf.TopPeersCount*4, n.conf.TopPeersCount*3)
}

func (n *Node) checkParents(client api.NodeClient, peer *Peer, parents hash.Events) {
	toDownload := n.lockNotDownloaded(parents)
	defer n.unlockDownloaded(toDownload)

	for e := range toDownload {
		var req api.EventRequest
		req.Hash = e.Bytes()

		n.downloadEvent(client, peer, &req)
	}
}

func (n *Node) compareKnownEvents(client api.NodeClient, peer *Peer) (map[hash.Peer]uint64, error) {
	knowns := n.knownEvents()

	req := &api.KnownEvents{
		Lasts: make(map[string]uint64, len(knowns)),
	}
	for id, h := range knowns {
		req.Lasts[id.Hex()] = h
	}

	ctx, cancel := context.WithTimeout(context.Background(), n.conf.ClientTimeout)
	defer cancel()

	id, ctx := api.ServerPeerID(ctx)

	resp, err := client.SyncEvents(ctx, req)
	if err != nil {
		n.ConnectFail(peer, err)
		return nil, err
	}

	if *id != peer.ID {
		// TODO: skip or continue gossiping with peer id ?
	}

	res := make(map[hash.Peer]uint64, len(resp.Lasts))
	for hex, h := range PeersHeightsDiff(resp.Lasts, req.Lasts) {
		res[hash.HexToPeer(hex)] = h
	}

	n.ConnectOK(peer)
	return res, nil
}

// downloadEvent downloads event.
func (n *Node) downloadEvent(client api.NodeClient, peer *Peer, req *api.EventRequest) (*inter.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), n.conf.ClientTimeout)
	defer cancel()

	id, ctx := api.ServerPeerID(ctx)

	w, err := client.GetEvent(ctx, req)
	if err != nil {
		n.ConnectFail(peer, err)
		return nil, err
	}

	if *id != peer.ID {
		// TODO: skip or continue gossiping with peer id ?
	}

	if w.Creator != req.PeerID || w.Index != req.Index {
		n.ConnectFail(peer, fmt.Errorf("bad GetEvent() response"))
		return nil, nil
	}

	event := inter.WireToEvent(w)

	// check event sign
	creator := n.store.GetPeer(event.Creator)
	if creator == nil {
		return nil, nil
	}
	if !event.Verify(creator.PubKey) {
		err = fmt.Errorf("falsity GetEvent() response")
		n.ConnectFail(peer, err)
		return nil, err
	}

	n.saveNewEvent(event)

	return event, nil
}

// saveNewEvent writes event to store, indexes and consensus.
// Note: event should be last from its creator.
func (n *Node) saveNewEvent(e *inter.Event) {
	n.store.SetEvent(e)
	n.store.SetEventHash(e.Creator, e.Index, e.Hash())
	n.store.SetPeerHeight(e.Creator, e.Index)

	if n.consensus != nil {
		n.consensus.PushEvent(e.Hash())
	}

	n.pushPotentialParent(e)
}

// knownEventsReq makes request struct with event heights of top peers.
func (n *Node) knownEvents() map[hash.Peer]uint64 {
	peers := n.peers.Snapshot()
	peers = append(peers, n.ID)

	res := make(map[hash.Peer]uint64, len(peers))
	for _, id := range peers {
		h := n.store.GetPeerHeight(id)
		res[id] = h
	}

	return res
}

/*
 * evaluation function for gossip
 */

// gossipEvaluation implements sort.Interface.
type gossipEvaluation Node

// Len is the number of elements in the collection.
func (n *gossipEvaluation) Len() int {
	return len(n.peers.top)
}

// Swap swaps the elements with indexes i and j.
func (n *gossipEvaluation) Swap(i, j int) {
	n.peers.top[i], n.peers.top[j] = n.peers.top[j], n.peers.top[i]
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (n *gossipEvaluation) Less(i, j int) bool {
	a := n.peers.attrByID(n.peers.top[i]).Host
	b := n.peers.attrByID(n.peers.top[j]).Host

	if a.LastSuccess.After(a.LastFail) && !b.LastSuccess.After(b.LastFail) {
		return true
	}

	if a.LastFail.After(a.LastSuccess) && b.LastFail.After(b.LastSuccess) {
		if a.LastFail.Before(b.LastFail) {
			return true
		}
	}

	if a.LastSuccess.After(b.LastSuccess) {
		return true
	}

	return false
}
