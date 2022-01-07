package config

import (
	"github.com/layer5io/meshery-adapter-library/adapter"
)

var (
	ServiceName = "service_name"
)

func getOperations(dev adapter.Operations) adapter.Operations {
	// var adapterVersions []adapter.Version
	// versions, _ := utils.GetLatestReleaseTagsSorted("meshery-perf", "meshery-perf")
	// for _, v := range versions {
	// 	adapterVersions = append(adapterVersions, adapter.Version(v))
	// }
	return dev
}
