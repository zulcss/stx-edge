package preflight

import (
	"os/exec"
	log "github.com/sirupsen/logrus"
)

func PreflightCheck() {
	// Check for Docker Cli
	DockerCheck()

	// Check for parted
	PartedCheck()

	// Check for luet
	LuetCheck()
}

// Check for Docker
func DockerCheck() {
	log.Info("Checking for Docker cli")

	path, err := exec.LookPath("docker")
	if err != nil {
		log.Fatal("Docker cli not found")
	}
	log.Infof("Docker CLI found: %s", path)
}

// Check for parted
func PartedCheck() {
	log.Info("Checking for parted")
	
	path, err := exec.LookPath("parted")
	if err != nil {
		log.Fatal("parted is not found")
	}
	log.Infof("Parted found: %s", path)
}

// Check for Luet
func LuetCheck() {
	path, err := exec.LookPath("luet")
	if err != nil {
		log.Fatal("luet is not found")
	}
	log.Infof("Luet found: %s", path)
}