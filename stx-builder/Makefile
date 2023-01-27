all: clean build

build:
	go build -o bin/stx-builder main.go

clean:
	rm -rf bin


dev:
	docker build -t stx-dev images/build
	docker run --privileged --group-add 0 --user 1000:1000 -v /dev:/dev -i -t -v /var/run/docker.sock:/var/run/docker.sock \
		-v $(PWD):/workdir -w /workdir stx-dev
