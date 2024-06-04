FROM golang:1.21-alpine as goBuilder
WORKDIR /arlog
COPY . /arlog/
RUN CGO_ENABLED=0 go build

FROM public.ecr.aws/zinclabs/rust:bookworm-sccache as builder
RUN rustc --version && sccache --version

WORKDIR /openobserve
COPY . /openobserve
COPY --from=webBuilder /web/dist web/dist

RUN cargo build --release

# FROM gcr.io/distroless/cc as runtime
FROM public.ecr.aws/debian/debian:bookworm-slim as runtime
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
RUN apt-get install -y curl htop iftop sysstat procps lsof net-tools
RUN update-ca-certificates
COPY --from=builder /openobserve/target/release/openobserve /
RUN ["/openobserve", "init-dir", "-p", "/data/"]
CMD ["/openobserve"]