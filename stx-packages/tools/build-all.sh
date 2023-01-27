#!/bin/bash

sudo luet build \
	-q \
	--full \
	--pull --image-repository local-repo --from-repositories \
	--no-spinner --live-output --tree packages
