FROM ubuntu:22.04

WORKDIR /MuxiFresh-Be-2.0/schedule

COPY rpc/schedule schedule-rpc
COPY rpc/etc/schedule.yaml etc/schedule-rpc.yaml

COPY api/schedule schedule-api
COPY api/etc/schedule-api.yaml etc/schedule-api.yaml

COPY start.sh ./

RUN chmod +x ./start.sh

CMD ["./start.sh"]