package disk

import (
	"io/ioutil"
	log "github.com/sirupsen/logrus"
)

type Disk struct {
	DiskName		string
	DiskSize	string
	DiskFS		string
	DiskType	string
	ImageTag	string
}

func NewDiskContext() *Disk {
	return &Disk {
		DiskName:	"disk.img",
		DiskSize:	"10G",
		DiskType:	"uefi",
		DiskFS:		"btrfs",
		ImageTag:	"stx-image:latest",
	}
}

func (d *Disk) CreateDisk() error {
	log.Debug("in CreateDisk")

	rootDir, _ := ioutil.TempDir("", "rootfs")
	log.Infof("Created rootfs: %s", rootDir)
	return nil
}

func IsValidDisk(diskfs string) bool {
	log.Debug("in IsValidDisk")
	switch diskfs {
	case
		"btrfs":
			return true
	}
	return false
}