package pipeline

import (
	"context"

	"golang.org/x/sync/errgroup"
)

//Stage interface
type Stage interface {
	Go(func() error)
	Done() <-chan struct{}
}

//Pipeline interface
type Pipeline interface {
	Stage() Stage
	Wait() error
}

type stage struct {
	g *errgroup.Group
	c context.Context
}

func (p *stage) Go(fn func() error) {
	p.g.Go(fn)
}

func (p *stage) Done() <-chan struct{} {
	return p.c.Done()
}

type pipeline struct {
	stage *stage
}

func (p *pipeline) Stage() Stage {
	return p.stage
}

func (p *pipeline) Wait() error {
	return p.stage.g.Wait()
}

//New creates a new pipline
func New() Pipeline {
	g, c := errgroup.WithContext(context.Background())
	return &pipeline{
		stage: &stage{
			g: g,
			c: c,
		},
	}
}
