package cmd

import (
	"os"
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
	"github.com/zulcss/stx-edge/pkg/images"
	"github.com/zulcss/stx-edge/pkg/utils"
)

var (
	PullImage	bool
	BuildImage	bool
	SourceDir	string
	ImageTag	string
)

var createImage = &cobra.Command {
	Use: "image",
	Short: "Create StartingX Container Image",
	Run: func(cmd *cobra.Command, args []string) {
		err := runCreate()
		if err != nil {
			log.Error(err)
		}
	},
}

func runCreate() error {
	log.Debugf("In runCrate")
	ctx := images.NewImageContext()
	if PullImage {
		log.Debug("in Pull image")
		if ImageTag == "" {
			log.Error("Image Tag not specified.")
			os.Exit(-1)
		} else {
			log.Infof("Pulling %s from registry.", ImageTag)
			ctx.ImageTag = ImageTag
			err := ctx.CreateImage()
			if err != nil {
				return nil
			}

		}
	}

	if BuildImage {
		log.Debug("in Build image")
		if ImageTag == "" {
			log.Error("Image Tag is not specified.")
			os.Exit(-1)
		}
		if SourceDir == "" {
			log.Error("Image source directory not specified.")
			os.Exit(-1)
		}

		if utils.PathExists(SourceDir) {
			log.Infof("Found image source: %s", SourceDir)
		}

		ctx.ImageTag = ImageTag
		ctx.ImageSourceDir = SourceDir
		err := ctx.BuildImage()
		if err != nil {
			os.Exit(-1)
		}
	}
	return nil
}


func init() {
	createImage.PersistentFlags().BoolVarP(&PullImage, "pull", "p", false, "Pull an image from a registry.")
	createImage.PersistentFlags().BoolVarP(&BuildImage, "build", "b", false, "Build from a local container.")
	createImage.PersistentFlags().StringVarP(&SourceDir, "source", "s", "", "Image source directory.")
	createImage.PersistentFlags().StringVarP(&ImageTag, "image", "i", "", "Image tag to pull.")
	createCommand.AddCommand(createImage)
}
