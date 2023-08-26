FROM centos:8.1.1911

WORKDIR /MuxiFresh-Be-2.0/userauth

#rpc
COPY cmd/rpc/accountCenter/accountCenter .
COPY cmd/rpc/accountCenter/etc/accountCenter.yaml etc/accountCenter.yaml

#api
COPY cmd/api/user-auth .
COPY cmd/api/etc/user-auth.yaml etc/user-auth.yaml

#redis
COPY ./redis-server .
COPY ./redis-cli .

#kafka
COPY opt/jdk /opt/jdk
COPY opt/zookeeper /opt/zookeeper
COPY opt/kafka /opt/kafka
COPY home/kafka-logs /home/kafka-logs
COPY etc/profile etc/profile

COPY start.sh ./

RUN chmod +x ./start.sh

CMD ["./start.sh"]