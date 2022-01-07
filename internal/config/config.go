package config

import (
	"path"

	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/common"
	"github.com/layer5io/meshery-adapter-library/config"
	"github.com/layer5io/meshery-adapter-library/status"
	configprovider "github.com/layer5io/meshkit/config/provider"
	"github.com/layer5io/meshkit/utils"
)

const (
	// Constants to use in log statements
	LabelNamespace = "label-namespace"

	ServicePatchFile = "service-patch-file"
	CPPatchFile      = "cp-patch-file"
	ControlPatchFile = "control-patch-file"
	FilterPatchFile  = "filter-patch-file"

	// Configure Envoy filter operation
	EnvoyFilterOperation = "envoy-filter-operation"

	// Addons that the adapter supports
	PrometheusAddon = "prometheus-addon"
	GrafanaAddon    = "grafana-addon"
	KialiAddon      = "kiali-addon"
	JaegerAddon     = "jaeger-addon"
	ZipkinAddon     = "zipkin-addon"

	// Policies
	DenyAllPolicyOperation     = "deny-all-policy-operation"
	StrictMTLSPolicyOperation  = "strict-mtls-policy-operation"
	MutualMTLSPolicyOperation  = "mutual-mtls-policy-operation"
	DisableMTLSPolicyOperation = "disable-mtls-policy-operation"

	// OAM Metadata constants
	OAMAdapterNameMetadataKey       = "adapter.meshery.io/name"
	OAMComponentCategoryMetadataKey = "ui.meshery.io/category"
)

var (
	MesheryPerfOperation = "meshery-perf"

	ServerVersion  = status.None
	ServerGitSHA   = status.None
	configRootPath = path.Join(utils.GetHome(), ".meshery")

	Config = configprovider.Options{
		FilePath: configRootPath,
		FileName: "meshery-perf",
		FileType: "yaml",
	}

	ServerConfig = map[string]string{
		"name":     "meshery-perf",
		"type":     "adapter",
		"port":     "9999",
		"traceurl": status.None,
	}

	MeshSpec = map[string]string{
		"name":    "meshery-perf",
		"status":  status.NotInstalled,
		"version": status.None,
	}

	ProviderConfig = map[string]string{
		configprovider.FilePath: configRootPath,
		configprovider.FileType: "yaml",
		configprovider.FileName: "meshery-perf",
	}

	// KubeConfig - Controlling the kubeconfig lifecycle with viper
	KubeConfig = map[string]string{
		configprovider.FilePath: configRootPath,
		configprovider.FileType: "yaml",
		configprovider.FileName: "kubeconfig",
	}

	Operations = getOperations(common.Operations)
)

// New creates a new config instance
func New(provider string) (h config.Handler, err error) {
	// Config provider
	switch provider {
	case configprovider.ViperKey:
		h, err = configprovider.NewViper(Config)
		if err != nil {
			return nil, err
		}
	case configprovider.InMemKey:
		h, err = configprovider.NewInMem(Config)
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrEmptyConfig
	}

	// Setup server config
	if err := h.SetObject(adapter.ServerKey, ServerConfig); err != nil {
		return nil, err
	}

	// Setup mesh config
	if err := h.SetObject(adapter.MeshSpecKey, MeshSpec); err != nil {
		return nil, err
	}

	// Setup Operations Config
	if err := h.SetObject(adapter.OperationsKey, Operations); err != nil {
		return nil, err
	}

	return h, nil
}

func NewKubeconfigBuilder(provider string) (config.Handler, error) {
	opts := configprovider.Options{
		FilePath: configRootPath,
		FileType: "yaml",
		FileName: "kubeconfig",
	}

	// Config provider
	switch provider {
	case configprovider.ViperKey:
		return configprovider.NewViper(opts)
	case configprovider.InMemKey:
		return configprovider.NewInMem(opts)
	}
	return nil, ErrEmptyConfig
}

// RootPath returns the config root path for the adapter
func RootPath() string {
	return configRootPath
}
