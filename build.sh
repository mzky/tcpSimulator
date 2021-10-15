#!/bin/bash
gox -osarch="linux/amd64" -ldflags "-w -s"
gox -osarch="linux/arm64" -ldflags "-w -s"
