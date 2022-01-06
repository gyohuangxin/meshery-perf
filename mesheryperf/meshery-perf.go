package mesheryperf

import (
	"github.com/layer5io/meshery-adapter-library/adapter"
	meshkitCfg "github.com/layer5io/meshkit/config"
	"github.com/layer5io/meshkit/logger"
)

// Meshery-perf represents the istio adapter and embeds adapter.Adapter
type MesheryPerf struct {
	adapter.Adapter // Type Embedded
}

// New initializes Meshery-perf handler.
func New(c meshkitCfg.Handler, l logger.Handler, kc meshkitCfg.Handler) adapter.Handler {
	return &MesheryPerf{
		Adapter: adapter.Adapter{
			Config:            c,
			Log:               l,
			KubeconfigHandler: kc,
		},
	}
}
