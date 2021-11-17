FROM ubuntu:bionic

RUN apt-get update && apt-get install -y build-essential autoconf automake gdb git libffi-dev zlib1g-dev libssl-dev curl vim sudo wget

RUN useradd -m -u 1000 -s /bin/bash -d /home/dylan dylan
RUN echo 'dylan ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

# Node 14, nodemon and golang
RUN set -x \
  && curl -fsSL https://deb.nodesource.com/setup_14.x | bash - \
  && apt-get install -y nodejs \
  && npm install -g nodemon \
  && rm -rf /usr/local/go \
  && wget https://golang.org/dl/go1.16.7.linux-amd64.tar.gz \
  && tar -C /usr/local -xzf go1.16.7.linux-amd64.tar.gz

USER dylan
ENV PATH=$PATH:/usr/local/go/bin