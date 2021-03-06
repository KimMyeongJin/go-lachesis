package proxy

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"github.com/Fantom-foundation/go-lachesis/src/common"
	"github.com/Fantom-foundation/go-lachesis/src/network"
	"github.com/Fantom-foundation/go-lachesis/src/poset"
	"github.com/Fantom-foundation/go-lachesis/src/proxy/proto"
)

func TestGrpcCalls(t *testing.T) {
	t.Run("over TCP", func(t *testing.T) {
		testGrpcCalls(t, network.TCPListener)
	})

	t.Run("over Fake", func(t *testing.T) {
		dialer := network.FakeDialer("client.fake")
		testGrpcCalls(t, network.FakeListener, grpc.WithContextDialer(dialer))
	})
}

func TestGrpcReConnection(t *testing.T) {
	t.Run("over TCP", func(t *testing.T) {
		testGrpcReConnection(t, network.TCPListener)
	})

	t.Run("over Fake", func(t *testing.T) {
		dialer := network.FakeDialer("client.fake")
		testGrpcReConnection(t, network.FakeListener, grpc.WithContextDialer(dialer))
	})
}

func testGrpcCalls(t *testing.T, listen network.ListenFunc, opts ...grpc.DialOption) {
	const (
		timeout    = 1 * time.Second
		errTimeout = "time is over"
	)
	var bind = "127.0.0.1:"

	logger := common.NewTestLogger(t)

	s, err := NewGrpcAppProxy(bind, timeout, logger, listen)
	if !assert.NoError(t, err) {
		return
	}
	bind = s.ListenAddr()

	c, err := NewGrpcLachesisProxy(bind, logger, opts...)
	if !assert.NoError(t, err) {
		return
	}

	t.Run("#1 Send tx", func(t *testing.T) {
		assertO := assert.New(t)
		gold := []byte("123456")

		err = c.SubmitTx(gold)
		assertO.NoError(err)

		select {
		case tx := <-s.SubmitCh():
			assertO.Equal(gold, tx)
		case <-time.After(timeout):
			assertO.Fail(errTimeout)
		}
	})

	t.Run("#2 Receive block", func(t *testing.T) {
		assert := assert.New(t)
		block := poset.Block{}
		gold := []byte("123456")

		go func() {
			select {
			case event := <-c.CommitCh():
				assert.Equal(block, event.Block)
				event.RespChan <- proto.CommitResponse{
					StateHash: gold,
					Error:     nil,
				}
			case <-time.After(timeout):
				assert.Fail(errTimeout)
			}
		}()

		answ, err := s.CommitBlock(block)
		if assert.NoError(err) {
			assert.Equal(gold, answ)
		}
	})

	t.Run("#3 Receive snapshot query", func(t *testing.T) {
		assert := assert.New(t)
		index := int64(1)
		gold := []byte("123456")

		go func() {
			select {
			case event := <-c.SnapshotRequestCh():
				assert.Equal(index, event.BlockIndex)
				event.RespChan <- proto.SnapshotResponse{
					Snapshot: gold,
					Error:    nil,
				}
			case <-time.After(timeout):
				assert.Fail(errTimeout)
			}
		}()

		answ, err := s.GetSnapshot(index)
		if assert.NoError(err) {
			assert.Equal(gold, answ)
		}
	})

	t.Run("#4 Receive restore command", func(t *testing.T) {
		assert := assert.New(t)
		gold := []byte("123456")

		go func() {
			select {
			case event := <-c.RestoreCh():
				assert.Equal(gold, event.Snapshot)
				event.RespChan <- proto.RestoreResponse{
					StateHash: gold,
					Error:     nil,
				}
			case <-time.After(timeout):
				assert.Fail(errTimeout)
			}
		}()

		err := s.Restore(gold)
		assert.NoError(err)
	})

	err = c.Close()
	assert.NoError(t, err)

	err = s.Close()
	assert.NoError(t, err)

}

func testGrpcReConnection(t *testing.T, listen network.ListenFunc, opts ...grpc.DialOption) {
	const (
		timeout    = 1 * time.Second
		errTimeout = "time is over"
	)
	var bind = "127.0.0.1:"

	logger := common.NewTestLogger(t)

	s, err := NewGrpcAppProxy(bind, timeout, logger, listen)
	if !assert.NoError(t, err) {
		return
	}
	bind = s.ListenAddr()

	c, err := NewGrpcLachesisProxy(bind, logger, opts...)
	if !assert.NoError(t, err) {
		return
	}
	if !assert.NotNil(t, c) {
		return
	}

	checkConnAndStopServer := func(t *testing.T) {
		assert := assert.New(t)
		gold := []byte("123456")

		err := c.SubmitTx(gold)
		assert.NoError(err)

		select {
		case tx := <-s.SubmitCh():
			assert.Equal(gold, tx)
		case <-time.After(timeout):
			assert.Fail(errTimeout)
		}

		err = s.Close()
		assert.NoError(err)
	}

	t.Run("#1 Send tx after connection", checkConnAndStopServer)

	s, err = NewGrpcAppProxy(bind, timeout/2, logger, listen)
	assert.NoError(t, err)

	<-time.After(timeout)

	t.Run("#2 Send tx after reconnection", checkConnAndStopServer)

	err = c.Close()
	assert.NoError(t, err)
}

/*
func TestGrpcMaxMsgSize(t *testing.T) {
	const (
		largeSize  = 100 * 1024 * 1024
		timeout    = 3 * time.Minute
		errTimeout = "time is over"

	)
	addr := utils.GetUnusedNetAddr(t);
	logger := common.NewTestLogger(t)

	s, err := NewGrpcAppProxy(addr, timeout, logger)
	assert.NoError(t, err)

	c, err := NewGrpcLachesisProxy(addr, logger)
	assert.NoError(t, err)

	largeData := make([]byte, largeSize)
	_, err = rand.Read(largeData)
	assert.NoError(t, err)

	t.Run("#1 Send large tx", func(t *testing.T) {
		assert := assert.New(t)

		err = c.SubmitTx(largeData)
		assert.NoError(err)

		select {
		case tx := <-s.SubmitCh():
			assert.Equal(largeData, tx)
		case <-time.After(timeout):
			assert.Fail(errTimeout)
		}
	})

	t.Run("#2 Receive large block", func(t *testing.T) {
		assert := assert.New(t)
		block := poset.Block{
			Body: poset.BlockBody{
				Transactions: [][]byte{
					largeData,
				},
			},
		}
		hash := largeData[:largeSize/10]

		go func() {
			select {
			case event := <-c.CommitCh():
				assert.EqualValues(block, event.Block)
				event.RespChan <- proto.CommitResponse{
					StateHash: hash,
					Error:     nil,
				}
			case <-time.After(timeout):
				assert.Fail(errTimeout)
			}
		}()

		answer, err := s.CommitBlock(block)
		if assert.NoError(err) {
			assert.Equal(hash, answer)
		}
	})

	err = c.Close()
	assert.NoError(t, err)

	err = s.Close()
	assert.NoError(t, err)
}
*/
