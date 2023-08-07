#!/bin/bash

nohup ./schedule-rpc -f etc/schedule-rpc.yaml &

./schedule-api -f etc/schedule-api.yaml
