#!/bin/bash

set -eux

export GOPATH="$(pwd)/.gobuild"
SRCDIR="${GOPATH}/src/github.com/RoyalIcing/plz"

[ -d ${GOPATH} ] && rm -rf ${GOPATH}
mkdir -p "${GOPATH}/src"
mkdir -p "${GOPATH}/pkg"
mkdir -p "${GOPATH}/bin"
mkdir -p "${SRCDIR}"
cp *.go "${SRCDIR}"
cp go.mod go.sum "${SRCDIR}"
cp -R cmd/ "${SRCDIR}/cmd"
(
    echo "${GOPATH}"
    cd "${SRCDIR}"
    go get .
    go install .
)
