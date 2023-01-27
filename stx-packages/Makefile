dev:
	docker build -t stx-packager
	docker run --privileged --group-add 0 --user 1000:1000 -v /dev:/dev -i -t -v /var/run/docker.sock:/var/run/docker.sock -v $(PWD):/workdir -w /workdir stx-build
