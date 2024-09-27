#!/bin/bash

BIN_PATH=$1

go clean

if [ -e "$BIN_PATH" ]; then 
  rm "$BIN_PATH"
fi