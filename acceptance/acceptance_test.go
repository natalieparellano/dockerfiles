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
		testDir, err := filepath.Abs(".")
		if err != nil {
			t.Fatal(err)
		}
		err = os.RemoveAll(filepath.Join(testDir, "in"))
		if err != nil {
			t.Fatal(err)
		}
		err = os.MkdirAll(filepath.Join(testDir, "in"), 0755)
		if err != nil {
			t.Fatal(err)
		}

		// build binary
		cmd := exec.Command("go", "build",
			"-o", filepath.Join(testDir, "testdata", "dockerfiles"),
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
			"-v", fmt.Sprintf("%s:/in", filepath.Join(testDir, "testdata")),
			"-v", fmt.Sprintf("%s:/kaniko", filepath.Join(testDir, "in")),
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
