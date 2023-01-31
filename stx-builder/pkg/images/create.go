package images

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/erikgeiser/promptkit/confirmation"

	"github.com/zulcss/stx-edge/pkg/utils"
)

type Image struct { 
	ImageTag	string
	ImageSourceDir	string
}

func NewImageContext() *Image {
	return &Image{
		ImageTag:	"stx-image:latest",
	}
}

// Pull an image from a Docker registry
func (i *Image) CreateImage() error {
	log.Debug("in CreateImage")

	exist := ImageExists(i.ImageTag)
	if exist {
		input := confirmation.New("Image already exists, do you want to update it?",
			confirmation.Undecided)
		update, _:= input.RunPrompt()
		if !update {
			return nil
		}
	}
	i.ImageFetch()

	return nil
}

// Build a local image from a dockerfile
func (i *Image) BuildImage() error {
	log.Debug("in BuildImage")

	err := i.ImageCreate()
	if err != nil {
		log.Error(err)
		return nil
	}
	
	return nil
}

// Pull a docker image from a registry
func (i *Image) ImageFetch() error {
	log.Debug("in FetchImage")

	out, err := utils.SH(fmt.Sprintf("docker image pull %s", i.ImageTag))
	if err != nil {
		log.Errorf("Failed to pull image: %s", out)
		return err
	}
	log.Infof("Succesfully pulled %s", i.ImageTag)
	return nil
}

func (i *Image) ImageCreate() error {
	log.Debugf("in ImageCreate")

	out, err := utils.SH(fmt.Sprintf("docker build -t %s %s", i.ImageTag, i.ImageSourceDir))
	if err != nil {
		log.Errorf("Failed to build container: %s", out)
		return err
	}
	log.Infof("Succesfully built %s: %s", i.ImageTag, i.ImageSourceDir)
	return nil
}

func ImageExists(image string) bool{
	log.Debug("in ImageExists")
	_, err := utils.SH(fmt.Sprintf("docker inspect --type=image %s", image))
	if err == nil {
	   log.Infof("Image is present locally: %s", image)
	   return true
	}
	return false
}
