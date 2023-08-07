#!/bin/bash

nohup ./redis-server &

chmod +x ./accountCenter
nohup ./accountCenter &

chmod +x ./user-auth
./user-auth
