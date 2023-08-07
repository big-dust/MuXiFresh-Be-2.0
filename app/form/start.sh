#!/bin/bash

nohup ./form-rpc -f etc/form-rpc.yaml &

./form-api -f etc/form-api.yaml

