package disk

import (
	"io/ioutil"
	
	log "github.com/sirupsen/logrus"

	"github.com/zulcss/stx-edge/pkg/images"
)

type Disk struct {
	DiskName		string
	DiskSize		string
	DiskFS			string
	DiskType		string
	ImageTag		string
	DiskRootDir		string
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
	d.DiskRootDir = rootDir
	log.Infof("Created rootfs: %s", d.DiskRootDir)

	log.Infof("Creating %s with %s", d.DiskName, d.DiskType)
	d.CreateDiskFormat()
	
	log.Infof("Creating filesystem %s on %s", d.DiskFS, d.DiskName)
	d.DiskFormat()

	log.Infof("Mounting %s on %s", d.DiskName, d.DiskRootDir)
	d.MountDisk()

	log.Infof("Unpacking %s on %s", d.ImageTag, d.DiskRootDir)
	images.Unpack(d.ImageTag, d.DiskRootDir)

	log.Infof("Unmounting %s from %s", d.DiskName, d.DiskRootDir)
	d.UmountDisk()
	return nil
}
