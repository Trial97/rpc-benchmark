package airc

import (
	"context"
)

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

type Arith int

// Some of Arith's methods have value args, some have pointer args. That's deliberate.

func (t *Arith) Add(ctx context.Context, args Args, reply *Reply) error {
	// select {
	// case <-ctx.Done():
	// return ctx.Err()
	// default:
	// }
	reply.C = args.A + args.B
	return nil
}
