package generator

import (
	"context"
	"errors"
	"sync/atomic"
)

const defaultBufferSize = 1000

var (
	ErrNotStarted     = errors.New("generator is not started")
	ErrAlreadyStarted = errors.New("generator is already started")
	ErrAlreadyStopped = errors.New("generator is already stopped")
)

type (
	Generator struct {
		stopCh        chan struct{}
		startedMarker int32
		stoppedMarker int32

		gen  Gen
		pipe chan interface{}
	}

	Gen interface {
		Value() interface{}
	}
)

func New(g Gen) *Generator {
	return NewWithBufferSize(g, defaultBufferSize)
}

func NewWithBufferSize(g Gen, bufferSize int) *Generator {
	return &Generator{
		stopCh:        make(chan struct{}),
		startedMarker: 0,
		stoppedMarker: 0,
		gen:           g,
		pipe:          make(chan interface{}, bufferSize),
	}
}

func (g *Generator) Start(ctx context.Context) error {
	if g.started() {
		return ErrAlreadyStarted
	}

	if g.stopped() {
		return ErrAlreadyStopped
	}

	atomic.AddInt32(&g.startedMarker, 1)
	go g.start(ctx)

	return nil
}

func (g *Generator) Stop() error {
	if g.stopped() {
		return ErrAlreadyStopped
	}

	if !g.started() {
		return ErrNotStarted
	}

	g.stop()

	return nil
}

func (g *Generator) Get() (interface{}, error) {
	if g.stopped() {
		return nil, ErrAlreadyStopped
	}

	if !g.started() {
		return nil, ErrNotStarted
	}

	return <-g.pipe, nil
}

func (g *Generator) start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			g.stop()
			return
		case <-g.stopCh:
			return
		default:
			g.pipe <- g.gen.Value()
		}
	}
}

func (g *Generator) started() bool {
	return atomic.LoadInt32(&g.startedMarker) > 0
}

func (g *Generator) stop() {
	atomic.AddInt32(&g.stoppedMarker, 1)
	atomic.StoreInt32(&g.startedMarker, 0)
	close(g.stopCh)
}

func (g *Generator) stopped() bool {
	return atomic.LoadInt32(&g.stoppedMarker) > 0
}
