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
	}
}
	

func (d *Disk) CreateRawDisk() error {
	out, err := utils.SH(fmt.Sprintf("truncate -s %s %s", d.DiskSize, d.DiskName))
	if err != nil {
		log.Errorf("Failed to create disk: %s", out)
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
		"raw":
			return true
	}
	return false
}
