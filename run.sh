#!/bin/bash
DIR_PREFIX=`pwd`
CONTAINER=stun_server
VERSION="1.0.0"
IMAGE=${CONTAINER}:${VERSION}
docker stop ${CONTAINER}
docker rm ${CONTAINER}
docker run -d --net=host \
              --name=${CONTAINER} \
              --env-file ${CONTAINER}.env \
	      --restart=always \
	      ${IMAGE}
              # tail -f /dev/null
