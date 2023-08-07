#!/bin/bash

nohup ./user-rpc -f etc/user-rpc.yaml &

./user-api -f etc/user-api.yaml
