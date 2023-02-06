ARG LUET_VERSION=0.33.0

FROM quay.io/luet/base:$LUET_VERSION as luet
FROM debian:testing-slim

ENV LUET_YES=yes
ENV LUET_NOLOCK=true

ADD overlay/ /
RUN apt-get update && \
    apt-get install -y ca-certificates \
	init \
	docker.io \
	procps \
	sudo \
	iproute2 \
	btrfs-progs \
	sytemd-container \
	systemd-boot \
	linux-image-amd64 \
	vim && \	
    rm -rf /var/lib/apt/lists/*

RUN useradd -m sysadmin && \
    echo "sysadmin:sysadmin" | chpasswd && \
    echo "root:root" | chpasswd

COPY --from=luet /usr/bin/luet /usr/bin/luet
RUN mkdir -p /etc/luet/repos.conf.d && \
  luet repo add stx-test -y --url quay.io/zulcss/stx-test --type docker && \
  luet install -y utils/yip
RUN echo "localhost" > /etc/hostname && \
    rm -f /etc/machine-id /var/lib/machine-id
