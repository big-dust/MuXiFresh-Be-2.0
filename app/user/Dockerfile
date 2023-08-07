FROM ubuntu:22.04

WORKDIR /MuxiFresh-Be-2.0/user

COPY cmd/rpc/user/user user-rpc
COPY cmd/rpc/user/etc/user.yaml etc/user-rpc.yaml

COPY cmd/api/user user-api
COPY cmd/api/etc/user.yaml etc/user-api.yaml

COPY start.sh ./

RUN chmod +x ./start.sh

CMD ["./start.sh"]