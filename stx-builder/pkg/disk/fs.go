package disk

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	
	"github.com/zulcss/stx-edge/pkg/utils"
	"github.com/zulcss/stx-edge/pkg/constants"
)

func (d *Disk) FormatExt() error {
	_, err := utils.SH(fmt.Sprintf("mkfs -L %s -t %s %s", constants.DiskLabel, d.DiskFS, d.DiskName))
	if err != nil {
		return err
	}
	return nil
}

func (d* Disk)DiskFormat() {
	log.Debug("in DiskFormat")
	switch d.DiskFS {
	case "ext4":
		d.FormatExt()
	}
}

