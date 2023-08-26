FROM centos:8.1.1911

WORKDIR /MuxiFresh-Be-2.0/task

COPY cmd/rpc/assignment/assignment assignment
COPY cmd/rpc/assignment/etc/assignment.yaml etc/assignment.yaml

COPY cmd/rpc/comment/comment comment
COPY cmd/rpc/comment/etc/comment.yaml etc/comment.yaml

COPY cmd/rpc/submission/submission submission
COPY cmd/rpc/submission/etc/submission.yaml etc/submission.yaml

COPY cmd/api/task task
COPY cmd/api/etc/task.yaml etc/task.yaml

COPY start.sh ./

RUN chmod +x ./start.sh

CMD ["./start.sh"]