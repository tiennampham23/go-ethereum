package native

import "github.com/ethereum/go-ethereum/eth/tracers"
import "github.com/tiennampham23/simulation-worker/tracer"

func init() {
	tracers.DefaultDirectory.Register(tracer.InternalTransactionTracerName, tracer.NewInternalTransactionTracer, false)
}
