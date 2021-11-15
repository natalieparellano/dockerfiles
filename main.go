package main

import (
	"fmt"
	"os"

	"github.com/GoogleContainerTools/kaniko/pkg/config"
	"github.com/GoogleContainerTools/kaniko/pkg/executor"
)

func main() {
	dockerfilePath := os.Args[1]
	if dockerfilePath == "" {
		panic("missing dockerfile path")
	}
	if _, err := os.Stat(dockerfilePath); err != nil {
		panic(err)
	}

	exportOption := os.Args[2]
	switch exportOption {
	case "local":
		exportLocal(dockerfilePath)
		return
	case "tarball":
		exportTarball(dockerfilePath)
		return
	default:
		panic(fmt.Sprintf("unsupported export option: %s", exportOption))
	}
}

func exportLocal(path string) {
	// TODO
}

func exportTarball(path string) {
	// create Kaniko config
	opts := &config.KanikoOptions{
		DockerfilePath: path,
		SnapshotMode:   "full",
	}

	// do build
	image, err := executor.DoBuild(opts)
	if err != nil {
		panic(err)
	}

	// get layers
	config, err := image.ConfigFile()
	if err != nil {
		panic(err)
	}
	fmt.Printf("config: %+v\n", config)
	layers, err := image.Layers()
	if err != nil {
		panic(err)
	}
	fmt.Printf("generated %d layers\n", len(layers))
	for _, layer := range layers {
		diffID, err := layer.DiffID()
		if err != nil {
			panic(err)
		}
		fmt.Printf("layer: %s\n", diffID)
	}
}
