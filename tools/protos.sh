#!/bin/bash

set -eu

cd src/altipla/protobuf/
protoc date.proto --go_out=$GOPATH/src
