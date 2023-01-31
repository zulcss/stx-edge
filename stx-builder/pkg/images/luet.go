package images

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	
	"github.com/zulcss/stx-edge/pkg/utils"
)


func Unpack(image string, rootdir string) error {
	log.Debug("in Unpack")

	_, err := utils.SH(fmt.Sprintf("luet util unpack --local %s %s", image, rootdir))
	if err != nil {
		return err
	}
	return nil
}