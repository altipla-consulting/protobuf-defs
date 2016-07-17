#!/bin/bash

set -eu

protoc ./ptypes/date.proto --go_out=$GOPATH/src
