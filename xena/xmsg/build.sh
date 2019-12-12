#!/bin/bash
set -e

rm -rf ./*pb.go
protoc --go_out=. *.proto
