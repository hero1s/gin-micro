#!/bin/sh

protoc --go_out=protos/user --proto_path=protos/user --micro_out=protos/user user.proto