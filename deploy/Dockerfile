FROM golang:1.21-alpine3.20 as goBuilder
USER root
WORKDIR /arlog
COPY . /arlog/

RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w --extldflags "-static -fpic"'

FROM alpine:3.20 as runtime
ENV LANG=C.UTF-8
ENV TZ=Asia/Shanghai
WORKDIR /
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
#    && apk add --no-cache vim iproute2 bash tini curl  tzdata \
#    && rm -rf /var/cache/apk/*
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk add  tzdata&& rm -rf /var/cache/apk/*
USER root
COPY --from=goBuilder /arlog/arlog /
COPY config.json  /
#CMD tini -g -- /arlog
CMD /arlog

# 打包命令  docker-slim build --http-probe=false --dockerfile deploy/Dockerfile --tag shanhy/spring-native-demo.slim:0.0.3 .
