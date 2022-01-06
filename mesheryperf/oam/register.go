package oam

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gyohuangxin/meshery-perf/internal/config"
	"github.com/layer5io/meshery-adapter-library/adapter"
)

var (
	basePath, _  = os.Getwd()
	workloadPath = filepath.Join(basePath, "templates", "oam", "workloads")
	traitPath    = filepath.Join(basePath, "templates", "oam", "traits")
	pathSets     = []schemaDefinitionPathSet{}
)

// AvailableVersions denote the component versions available statically
var AvailableVersions = map[string]bool{}

type schemaDefinitionPathSet struct {
	oamDefinitionPath string
	jsonSchemaPath    string
	name              string
}

// RegisterWorkloads will register all of the workload definitions
// present in the path oam/workloads
//
// Registration process will send POST request to $runtime/api/oam/workload
func RegisterWorkloads(runtime, host string) error {
	oamRDP := []adapter.OAMRegistrantDefinitionPath{}

	for _, pathSet := range pathSets {
		metadata := map[string]string{
			config.OAMAdapterNameMetadataKey: config.MesheryPerfOperation,
		}

		if strings.HasSuffix(pathSet.name, "addon") {
			metadata[config.OAMComponentCategoryMetadataKey] = "addon"
		}

		oamRDP = append(oamRDP, adapter.OAMRegistrantDefinitionPath{
			OAMDefintionPath: pathSet.oamDefinitionPath,
			OAMRefSchemaPath: pathSet.jsonSchemaPath,
			Host:             host,
			Metadata:         metadata,
		})
	}

	return adapter.
		NewOAMRegistrant(oamRDP, fmt.Sprintf("%s/api/oam/workload", runtime)).
		Register()
}

// RegisterTraits will register all of the trait definitions
// present in the path oam/traits
//
// Registeration process will send POST request to $runtime/api/oam/trait
func RegisterTraits(runtime, host string) error {
	oamRDP := []adapter.OAMRegistrantDefinitionPath{}

	pathSets, err := load(traitPath)
	if err != nil {
		return err
	}

	for _, pathSet := range pathSets {
		metadata := map[string]string{
			config.OAMAdapterNameMetadataKey: config.MesheryPerfOperation,
		}

		oamRDP = append(oamRDP, adapter.OAMRegistrantDefinitionPath{
			OAMDefintionPath: pathSet.oamDefinitionPath,
			OAMRefSchemaPath: pathSet.jsonSchemaPath,
			Host:             host,
			Metadata:         metadata,
		})
	}

	return adapter.
		NewOAMRegistrant(oamRDP, fmt.Sprintf("%s/api/oam/trait", runtime)).
		Register()
}

func load(basePath string) ([]schemaDefinitionPathSet, error) {
	res := []schemaDefinitionPathSet{}

	if err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if matched, err := filepath.Match("*_definition.json", filepath.Base(path)); err != nil {
			return err
		} else if matched {
			nameWithPath := strings.TrimSuffix(path, "_definition.json")

			res = append(res, schemaDefinitionPathSet{
				oamDefinitionPath: path,
				jsonSchemaPath:    fmt.Sprintf("%s.meshery.layer5io.schema.json", nameWithPath),
				name:              filepath.Base(nameWithPath),
			})
			AvailableVersions[filepath.Base(filepath.Dir(path))] = true

		}

		return nil
	}); err != nil {
		return nil, err
	}

	return res, nil
}
func init() {
	var err error
	pathSets, err = load(workloadPath)
	if err != nil {
		fmt.Printf("Could not load definitions and schemas for static component registeration: %v", err.Error())
		return
	}
}
