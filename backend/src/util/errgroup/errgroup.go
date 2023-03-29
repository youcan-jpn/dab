package errgroup

import (
	"context"
	"fmt"
	"sync"

	eg "golang.org/x/sync/errgroup"
)

type Group struct {
	gr *eg.Group
	o  sync.Once
}

func WithContext(ctx context.Context) (*Group, context.Context) {
	g, ctx := eg.WithContext(ctx)
	return &Group{gr: g}, ctx
}

func (g *Group) Wait() error {
	if g.gr == nil {
		return nil
	}
	return g.gr.Wait()
}

func (g *Group) Go(ctx context.Context, f func() error) {
	if g.gr == nil {
		g.o.Do(func() {
			g.gr = &eg.Group{}
		})
	}
	g.gr.Go(func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("recovered from panic: %v", r)
			}
		}()
		return f()
	})
}
