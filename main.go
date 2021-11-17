package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/GoogleContainerTools/kaniko/pkg/config"
	"github.com/GoogleContainerTools/kaniko/pkg/executor"
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

const ( // TODO: derive or pass in
	outputDir                  = "/kaniko"
	ubuntuBionicTopLayerDigest = "sha256:284055322776031bac33723839acb0db2d063a525ba4fa1fd268a831c7553b26"
)

func main() {
	exportTarball()
}

func exportTarball() {
	// create Kaniko config
	opts := &config.KanikoOptions{
		CacheOptions:   config.CacheOptions{CacheDir: "/cache"},
		DockerfilePath: "/workspace/Dockerfile",
		IgnoreVarRun:   true,
		NoPush:         true,
		SrcContext:     "dir:///workspace/",
		SnapshotMode:   "full",
	}

	if err := os.Chdir("/"); err != nil {
		panic(err)
	}

	// do build
	image, err := executor.DoBuild(opts)
	if err != nil {
		panic(err)
	}

	// get config file
	configJSON, err := image.ConfigFile()
	if err != nil {
		panic(err)
	}
	configPath := filepath.Join(outputDir, "config.json")
	c, err := os.Create(configPath)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	err = json.NewEncoder(c).Encode(*configJSON)

	// get layers
	layers, err := image.Layers()
	if err != nil {
		panic(err)
	}
	fmt.Printf("generated %d layers\n", len(layers))
	for _, layer := range layers {
		digest, err := layer.Digest()
		if err != nil {
			panic(err)
		}
		digestStr := digest.String()
		if digestStr == ubuntuBionicTopLayerDigest {
			continue
		}
		layerPath := filepath.Join(outputDir, digestStr+".tgz")
		err = saveLayer(layer, layerPath)
		if err != nil {
			panic(err)
		}
	}
}

func saveLayer(layer v1.Layer, path string) error {
	layerReader, err := layer.Compressed()
	if err != nil {
		return err
	}
	l, err := os.Create(path)
	if err != nil {
		return err
	}
	defer l.Close()
	_, err = io.Copy(l, layerReader)
	if err != nil {
		return err
	}
	return nil
}
