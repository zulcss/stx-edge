#!/bin/bash
#
sudo luet create-repo \
		 --name "stx-test" \
                --type docker \
                --output "quay.io/zulcss/stx-test" \
                --push-images \
                --packages build \
                --force-push \
                --tree packages
