FROM golang:1.21-alpine as goBuilder
WORKDIR /arlog
COPY . /arlog/
RUN CGO_ENABLED=0 go build

FROM alpine:3.18.6 as runtime

COPY --from=goBuilder /arlog/arlog /
CMD /arlog