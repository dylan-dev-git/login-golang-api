FROM ubuntu:bionic

RUN apt-get update && apt-get install -y build-essential autoconf automake gdb git libffi-dev zlib1g-dev libssl-dev curl vim sudo

RUN useradd -m -u 1000 -s /bin/bash -d /home/dylan dylan
RUN echo 'dylan ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

# Node 14
RUN set -x \
  && curl -fsSL https://deb.nodesource.com/setup_14.x | bash - \
  && apt-get install -y nodejs

# Angular 12
RUN npm install -g @angular/cli@12.0.3

USER dylan