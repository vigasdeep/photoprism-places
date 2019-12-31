#!/usr/bin/env bash

PLACES_DATE=`date -u +%y%m%d`
PLACES_VERSION=`git describe --always`

if [[ -z $1 ]] || [[ -z $2 ]]; then
    echo "Please provide build mode and output file name" 1>&2
    exit 1
fi

if [[ $OS == "Windows_NT" ]]; then
    PLACES_OS=win32
    if [[ $PROCESSOR_ARCHITEW6432 == "AMD64" ]]; then
        PLACES_ARCH=amd64
    else
        if [[ $PROCESSOR_ARCHITECTURE == "AMD64" ]]; then
            PLACES_ARCH=amd64
        fi
        if [[ $PROCESSOR_ARCHITECTURE == "x86" ]]; then
            PLACES_ARCH=ia32
        fi
    fi
else
   PLACES_OS=`uname -s`
   PLACES_ARCH=`uname -p`
fi

if [[ $1 == "debug" ]]; then
  echo "Building development binary..."
	go build -ldflags "-X main.version=${PLACES_DATE}-${PLACES_VERSION}-${PLACES_OS}-${PLACES_ARCH}-DEBUG" -o $2 cmd/places/places.go
	du -h $2
	echo "Done."
elif [[ $1 == "static" ]]; then
  echo "Building static production binary..."
	go build -a -v -ldflags "-linkmode external -extldflags \"-static -L /usr/lib\" -s -w -X main.version=${PLACES_DATE}-${PLACES_VERSION}-${PLACES_OS}-${PLACES_ARCH}" -o $2 cmd/places/places.go
	du -h $2
	echo "Done."
else
  echo "Building production binary..."
	go build -ldflags "-s -w -X main.version=${PLACES_DATE}-${PLACES_VERSION}-${PLACES_OS}-${PLACES_ARCH}" -o $2 cmd/places/places.go
	du -h $2
	echo "Done."
fi
