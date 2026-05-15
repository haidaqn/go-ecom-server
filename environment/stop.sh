#!/bin/bash

# Start containers
docker-compose -f environment/docker-compose-dev.yml down
echo "[DANG]: vetautet server stop..."