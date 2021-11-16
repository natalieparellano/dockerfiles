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

		fmt.Println("building binary...")
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

		fmt.Println("executing binary in container...")
		cmd = exec.Command("docker", "run",
			"--rm",
			"--env", "PATH=/usr/local/bin:/kaniko",
			"--env", "HOME=/root",
			"--env", "USER=root",
			"--env", "SSL_CERT_DIR=/kaniko/ssl/certs",
			"--env", "DOCKER_CONFIG=/kaniko/.docker/",
			"--env", "DOCKER_CREDENTIAL_GCR_CONFIG=/kaniko/.config/gcloud/docker_credential_gcr_config.json",
			"-v", fmt.Sprintf("%s:/workspace", filepath.Join(testDir, "testdata")),
			"-v", fmt.Sprintf("%s:/kaniko", filepath.Join(testDir, "in")),
			"golang",
			"/workspace/dockerfiles", "/workspace/Dockerfile", "tarball",
		)
		output, err = cmd.CombinedOutput()
		fmt.Println(string(output))
		if err != nil {
			t.Fatal(err)
		}
	})
}
