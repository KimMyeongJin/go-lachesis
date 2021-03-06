package proxy

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/Fantom-foundation/go-lachesis/src/hash"
	"github.com/Fantom-foundation/go-lachesis/src/inter"
	"github.com/Fantom-foundation/go-lachesis/src/network"
	"github.com/Fantom-foundation/go-lachesis/src/proxy/wire"
)

var (
	// ErrConnTimeout returns when deadline exceeded
	// for Ctrl server connection.
	ErrConnTimeout = errors.New("node connection timeout")
)

// Node representation.
type Node interface {
	GetID() hash.Peer
	AddInternalTxn(inter.InternalTransaction)
}

// Consensus representation.
type Consensus interface {
	GetStakeOf(peer hash.Peer) float64
}

// CtrlAppProxy handles managing requests.
type CtrlAppProxy struct {
	Node
	Consensus
}

// InternalTxn pushes internal transaction into the Node.
func (s *CtrlAppProxy) InternalTxn(ctx context.Context, req *wire.InternalTxnRequest) (*empty.Empty, error) {
	peer := hash.HexToPeer(req.Receiver)

	tx := inter.InternalTransaction{
		Amount:   req.Amount,
		Receiver: peer,
	}

	s.Node.AddInternalTxn(tx)

	resp := empty.Empty{}
	return &resp, nil
}

// Stake returns the Node stake.
func (s *CtrlAppProxy) Stake(ctx context.Context, _ *empty.Empty) (*wire.StakeResponse, error) {
	peer := s.Node.GetID()
	resp := wire.StakeResponse{
		Value: s.Consensus.GetStakeOf(peer),
	}

	return &resp, nil
}

// ID returns the Node id.
func (s *CtrlAppProxy) ID(ctx context.Context, _ *empty.Empty) (*wire.IDResponse, error) {
	peer := s.Node.GetID()
	resp := wire.IDResponse{
		Id: peer.Hex(),
	}

	return &resp, nil
}

// NewCtrlAppProxy starts Ctrl server.
func NewCtrlAppProxy(bindAddr string, n Node, c Consensus, listen network.ListenFunc) *grpc.Server {
	srv := CtrlAppProxy{
		Node:      n,
		Consensus: c,
	}
	s := grpc.NewServer()
	wire.RegisterCtrlServer(s, &srv)

	if listen == nil {
		listen = network.TCPListener
	}
	listener := listen(bindAddr)

	go func() {
		if err := s.Serve(listener); err != nil {
			panic(err)
		}
	}()

	return s
}

// NewCtrlClient returns client for lachesis Ctrl.
func NewCtrlClient(addr string, timeout time.Duration) (wire.CtrlClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		if errors.Cause(err) == context.DeadlineExceeded {
			return nil, ErrConnTimeout
		}
		return nil, err
	}

	cli := wire.NewCtrlClient(conn)
	return cli, nil
}
