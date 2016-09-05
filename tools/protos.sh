#!/bin/bash

set -eu

function build {
  protoc -I $GOPATH/src $(pwd)/$1 --go_out=$GOPATH/src
}

build api/acl/acl.proto
build api/annotations/annotations.proto
build date/date.proto
