#!/bin/sh

pushd _script
go run ./histver -file ../users/releases.csv
