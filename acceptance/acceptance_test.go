package acceptance

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestTarball(t *testing.T) {
	t.Run("tarball", func(t *testing.T) {
		// setup
		srcPath, err := filepath.Abs(".") // acceptance directory
		if err != nil {
			t.Fatal(err)
		}
		err = os.RemoveAll(filepath.Join(srcPath, "in"))
		if err != nil {
			t.Fatal(err)
		}
		err = os.MkdirAll(filepath.Join(srcPath, "in"), 0755)
		if err != nil {
			t.Fatal(err)
		}

		// build binary
		cmd := exec.Command("go", "build",
			"-o", filepath.Join(srcPath, "testdata", "dockerfiles"),
			"..",
		)
		cmd.Env = append(os.Environ(), "GOOS=linux")
		output, err := cmd.CombinedOutput()
		fmt.Println(string(output))
		if err != nil {
			t.Fatal(err)
		}

		// execute binary in container
		cmd = exec.Command("docker", "run",
			"--rm",
			"-v", fmt.Sprintf("%s:/in", filepath.Join(srcPath, "testdata")),
			"-v", fmt.Sprintf("%s:/kaniko", filepath.Join(srcPath, "in")),
			"golang",
			"/in/dockerfiles", "/in/Dockerfile", "tarball",
		)
		output, err = cmd.CombinedOutput()
		fmt.Println(string(output))
		if err != nil {
			t.Fatal(err)
		}
	})
}
