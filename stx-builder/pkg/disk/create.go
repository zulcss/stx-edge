package disk

import (
	log "github.com/sirupsen/logsrus"
)

type Disk struct {
	DiskTag		string
	DiskSize	string
	DiskType	string
	DiskFS		string
}

func NewDiskContext() *Disk {
	return &Disk {
		DiskTag:	"disk.img",
		DiskSize	"10G",
		DiskType	"uefi",
		DiskFS		"btrfs"
	}
}

func (d *Disk) CreateDisk() error {
	log.Debug("in CreateDisk")
	return nil
}
