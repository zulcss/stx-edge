ARG LUET_VERSION=0.33.0

FROM quay.io/luet/base:$LUET_VERSION as luet
FROM debian:testing

ENV LUET_YES=yes
ENV LUET_NOLOCK=true

RUN apt-get update && \
    apt-get install -y ca-certificates \
	docker.io \
	golang \
	build-essential \
	parted \
	curl \
	btrfs-progs \
	vim \
	sudo && \
    rm -rf /var/lib/apt/lists/*

RUN adduser dev  && usermod -aG docker dev
RUN echo "dev:dev" | chpasswd
RUN echo '%dev ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers
RUN chmod a+s $(which docker)

COPY --from=luet /usr/bin/luet /usr/bin/luet
RUN mkdir -p /etc/luet/repos.conf.d && \
  luet repo add stx-test -y --url quay.io/zulcss/stx-test --type docker
RUN luet install -y utils/yip
