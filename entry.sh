#!/bin/bash


echo "IP[${PUBLIC_IP}] PORT[${STUN_PORT}]"
if [ "$1" = "" ]; then
	CMD="/main -port=${STUN_PORT} -public-ip=${PUBLIC_IP}"
else
        CMD="$*"
fi

echo "Running [$CMD]"
exec $CMD
echo "exiting ..."
