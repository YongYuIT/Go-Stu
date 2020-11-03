#!/bin/bash

dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient /go/src/main.go
