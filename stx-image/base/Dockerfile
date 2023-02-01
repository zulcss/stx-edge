ARG LUET_VERSION=0.33.0

FROM quay.io/luet/base:$LUET_VERSION as luet
FROM debian:testing-slim

ENV LUET_YES=yes
ENV LUET_NOLOCK=true

ADD overlay/ /
RUN apt-get update && \
    apt-get install -y ca-certificates \
	init \
	procps \
	sudo \
	iproute2 \
	openssh-server \
	btrfs-progs \
	nano \
	jq \
	vim && \	
    rm -rf /var/lib/apt/lists/*

RUN useradd -m sysadmin && \
    echo "sysadmin:sysadmin" | chpasswd

COPY --from=luet /usr/bin/luet /usr/bin/luet
RUN mkdir -p /etc/luet/repos.conf.d && \
  luet repo add stx-test -y --url quay.io/zulcss/stx-test --type docker && \
  luet install -y utils/yip
