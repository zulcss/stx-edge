#!/bin/bash

#sudo luet build \
#	-q \
#	--pull --image-repository local-repo --from-repositories \
#	--no-spinner --live-output --tree packages $1

sudo luet build --all --destination build --tree packages \
		--compression gzip $1

