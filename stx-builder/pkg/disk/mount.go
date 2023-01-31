package disk

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	
	"github.com/zulcss/stx-edge/pkg/utils"
)

func (d *Disk) MountDisk() error {
	log.Debug("in MountDisk")
	_, err := utils.SH(fmt.Sprintf("mount -t %s %s %s", d.DiskFS, d.DiskName, d.DiskRootDir))
	if err != nil {
		log.Errorf("Failed to mount %s on %s", d.DiskName, d.DiskRootDir)
		return err
	}
	return nil
}

func (d *Disk) UmountDisk() error {
	log.Debug("in UmountDisk")
	_, err := utils.SH(fmt.Sprintf("umount %s", d.DiskRootDir))
	if err != nil {
		log.Errorf("Failed to umount %s", d.DiskRootDir)
		return err
	}
	return nil
}