package cmd

import (
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
)

var (
	// Name of the disk image
	DiskName	string
	// Size of the disk
	DiskSize	string
	// Disk format (ext4, btrfs, etc)
	DiskFS		string
	// Type of disk (bare, uefi, qcow2, etc)
	DiskType	string
	// Directory to store the artifact
	OutputDir	string
	// Image to unpack
	ImgeTag		string
)

var createDisk = &cobra.Command {
	Use: "disk",
	Short: "Create StarlingX Disk",
	Run: func(cmd *cobra.Command, args[] string) {
		err := CreateDisk()
		if err != nil {
			log.Error(err)
		}
	},
}

func CreateDisk() error {
	return nil
}

func init() {
	createDisk.PersistentFlags().StringVarP(&DiskName, "disk", "d", "", "Name of Disk.")
	createDisk.PersistentFlags().StringVarP(&DiskSize, "size", "s", "", "Size of Disk.")
	createDisk.PersistentFlags().StringVarP(&DiskFS, "format", "f", "btrfs", "Filesystem format.")
	createDisk.PersistentFlags().StringVarP(&DiskType, "type", "t", "uefi", "Type of Disk.")
	createDisk.PersistentFlags().StringVarP(&OutputDir, "output", "o", "output", "Disk artifact directory")
	createDisk.PersistentFlags().StringVarP(&ImageTag, "image", "i", "stx-image", "Image to Unpack")
	createCommand.AddCommand(createDisk)
}

