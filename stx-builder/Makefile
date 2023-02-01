all: clean build

build:
	go build -o bin/stx-builder main.go

clean:
	rm -rf bin

disk:
	rm -f disk.img
	sudo ./bin/stx-builder create disk \
		--image stx-image \
		--output disk.img \
		--size 20G \
		--format ext4 \
		--type uefi 
