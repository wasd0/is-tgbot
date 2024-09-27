#!/bin/bash

BIN_PATH=$1

if [ -e "$BIN_PATH" ]; then
  $BIN_PATH
else
  make build
  $BIN_PATH
fi