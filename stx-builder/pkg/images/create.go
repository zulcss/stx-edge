package images

import (
	"fmt"
	"github.com/zulcss/stx-edge/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type Image struct { 
	ImageTag	string
}

func NewImageContext() *Image {
	return &Image{
		ImageTag:	"stx-image:latest",
	}
}

// Pull an image from a Docker registry
func (i *Image) CreateImage() error {
	log.Debug("in CreateImage")
	i.FetchImage()

	return nil
}

// Pull a docker image from a registry
func (i *Image) FetchImage() error {
	log.Debug("in FetchImage")

	out, err := utils.SH(fmt.Sprintf("docker image pull %s", i.ImageTag))
	if err != nil {
		log.Errorf("Failed to pull image: %s", out)
		return err
	}
	return nil
}
