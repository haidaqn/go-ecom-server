#!/bin/bash

#docker_build
make docker_build

#
echo "127.0.0.1 goecmbackendapi.com" >> /etc/hosts
