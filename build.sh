#!/bin/bash
CURRENT_COMMIT=$1
CONTAINER="stun_server"
VERSION="1.0.0"


docker build . -f Dockerfile -t ${CONTAINER}:${VERSION}
