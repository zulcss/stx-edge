FROM debian:testing

RUN apt-get update && \
    apt-get install -y ca-certificates \
	docker.io \
	golang \
	build-essential \
	sudo \
	vim 

RUN adduser dev  && usermod -aG docker dev
RUN echo "dev:dev" | chpasswd
RUN echo '%dev ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers
RUN chmod a+s $(which docker)

USER dev
