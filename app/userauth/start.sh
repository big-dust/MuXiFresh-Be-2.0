#!/bin/bash

source /etc/profile

cd /opt/zookeeper/bin

zkServer.sh start

cd ../../kafka

nohup bin/kafka-server-start.sh config/server.properties &

cd /MuxiFresh-Be-2.0/userauth

nohup ./redis-server &

chmod +x ./accountCenter

nohup ./accountCenter &

chmod +x ./user-auth

./user-auth
