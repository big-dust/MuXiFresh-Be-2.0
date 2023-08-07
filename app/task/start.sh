#!/bin/bash

nohup ./assignment -f etc/assignment.yaml &

nohup ./comment -f etc/comment.yaml &

nohup ./submission -f etc/submission.yaml &

./task -f etc/task.yaml
