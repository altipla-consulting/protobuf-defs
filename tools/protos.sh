#!/bin/bash

set -eu

protoc date/date.proto --go_out=$GOPATH/src
