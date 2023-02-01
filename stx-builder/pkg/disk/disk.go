package disk

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/zulcss/stx-edge/pkg/utils"
)

func(d *Disk) CreateDiskFormat() {
	log.Debug("in CreateDiskFormat")

	switch d.DiskType {
	case "raw":
		d.CreateRawDisk()
	case "gpt":
		d.CreateEFIDisk()
	}
}
	
func (d *Disk) CreateRawDisk() error {
	log.Debug("in CreateRawDisk")
	_, err := utils.SH(fmt.Sprintf("truncate -s %s %s", d.DiskSize, d.DiskName))
	if err != nil {
		log.Errorf("Failed to create disk: %s", err)
		return err
	}
	return nil
}

func (d *Disk) CreateGPTisk() error {
	log.Debug("in CreateGPTDisk")

	// Create a raw disk before running parted
	d.CreateRawDisk()

	out, err := utils.SH(fmt.Sprintf("parted -s -a optimal %s mklabel gpt", d.DiskName))
	if err != nil {
		log.Errorf("Failed to create disk: %s", out)
		return err
	}
	return nil
}

func IsValidDisk(diskfs string) bool {
	log.Debug("in IsValidDisk")
	switch diskfs {
	case
		"ext4":
			return true
	}
	return false
}

func IsValidDiskFormat(format string) bool {
	log.Debug("in IsValidDiskFormat")
	switch format {
	case
		"raw", "gpt":
			return true
	}
	return false
}
