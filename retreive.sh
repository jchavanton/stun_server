#!/bin/bash

INSTALL_PREFIX="/opt/halo"

# declare -a stun_server_instances=("hostname1" "hostname2")
source hosts.sh

retreive_stun_server_config() {
	ROLE="stun_server"
	INSTALL_DIR="${INSTALL_PREFIX}/${ROLE}"
	for i in "${stun_server_instances[@]}"
	do
		if [ "$1" != "all" ] && [ "$1" != "$i" ] ; then continue; fi
		printf "\ndownloading from [$i]\n"
		scp $i:$INSTALL_DIR/* .
		done
}

instruction() {
	printf  "\nYou can specify a host name :\n\n"
	for i in "${stun_server_instances[@]}"
	do
		echo "./retreive.sh $i"
	done
}

TARGET="$1"
if [ "${TARGET}" == "" ]
then
	instruction
	exit
fi

retreive_stun_server_config $TARGET
