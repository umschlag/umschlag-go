#!/usr/bin/env bash
set -eo pipefail

if [ -L $0 ]
then
    ROOT=$(realpath -e $(dirname $(readlink -e $0))/..)
else
    ROOT=$(realpath -e $(dirname $0)/..)
fi

SPEC=${SPEC:-https://dl.umschlag.tech/openapi/1.0.0-alpha1.yml}

docker run --rm \
	-v ${ROOT}:/generate \
	toolhippie/goswagger:latest \
	generate \
	client \
	-f ${SPEC} \
	-t /generate \
	-c umschlag
