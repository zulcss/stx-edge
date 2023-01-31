package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
	
	"github.com/zulcss/stx-edge/pkg/disk"
	"github.com/zulcss/stx-edge/pkg/utils"
)

var (
	// Name of the disk image
	Output	string
	// Size of the disk
	DiskSize	string
	// Disk format (ext4, btrfs, etc)
	DiskFS		string
	// Type of disk (bare, uefi, qcow2, etc)
	DiskType	string
	// Image to unpack
	ImgeTag		string
)

var createDisk = &cobra.Command {
	Use: "disk",
	Short: "Create StarlingX Disk",
	Run: func(cmd *cobra.Command, args[] string) {
		// Must run as root because of filesystem operations
		if !utils.CheckUser() {
			log.Fatal("This command requires root to be run.")
			os.Exit(-1)
		}
	  	err := InitDisk()
		if err != nil {
			log.Error(err)
			os.Exit(-1)
		}
	},
}

// Check for arguments and create the disk
func InitDisk() error {
	log.Debug("in CreateDisk")

	ctx := disk.NewDiskContext()
	if Output  == "" {
		log.Fatal("You did not specify a disk.")
	}

	if utils.PathExists(Output) {
		log.Fatalf("%s already exists. Please remove it before running this command ", Output)
		os.Exit(-1)
	}
	path, _ := filepath.Abs(Output)
	ctx.DiskName = path
	log.Infof("Using %s for disk path", ctx.DiskName)

	log.Infof("Checking for valid filesystem: %s", DiskFS)
	if DiskFS == "" {
		log.Fatalf("You did not specify a filesystem.")
	}
	if !disk.IsValidDisk(DiskFS) {
		log.Fatalf("%s is not a valid filesystem", DiskFS)
		os.Exit(-1)
	}
	/*ctx.DiskTag = DiskName
	err := ctx.CreateDisk()
	*/

	return nil
}

func init() {
	createDisk.PersistentFlags().StringVarP(&Output, "output", "o", "", "Name of disk image.")
	createDisk.PersistentFlags().StringVarP(&DiskSize, "size", "s", "", "Size of disk image.")
	createDisk.PersistentFlags().StringVarP(&DiskFS, "format", "f", "", "Filesystem type of disk image.")
	createDisk.PersistentFlags().StringVarP(&DiskType, "type", "t", "uefi", "Type of disk image.")
	createDisk.PersistentFlags().StringVarP(&ImageTag, "image", "i", "stx-image", "Image to unpack")
	createCommand.AddCommand(createDisk)
}

