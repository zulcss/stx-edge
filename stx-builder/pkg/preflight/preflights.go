package preflight

import (
	"os/exec"
	log "github.com/sirupsen/logrus"
)

func PreflightCheck() {
	// Check for Docker Cli
	DockerCheck()
}

func DockerCheck() {
	log.Info("Checking for Docker cli")

	path, err := exec.LookPath("docker")
	if err != nil {
		log.Error("Docker cli not found")
	}
	log.Infof("Docker CLI found: %s", path)
}
