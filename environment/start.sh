#!/bin/bash

# Start containers
docker-compose -f environment/docker-compose-dev.yml up
echo "[DANG]: vetautet server start..."