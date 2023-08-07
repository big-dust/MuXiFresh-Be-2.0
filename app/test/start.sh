#!/bin/bash

nohup ./test-rpc -f etc/test-rpc.yaml &

./test-api -f etc/test-api.yaml

