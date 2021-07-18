
FROM golang as builder
RUN apt-get update && apt-get install -y cmake build-essential
WORKDIR /build
COPY . ./
RUN git submodule update --init --recursive && make && C_INCLUDE_PATH=/build/go-llama.cpp LIBRARY_PATH=/build/go-llama.cpp go build -o app .

FROM debian:12
RUN apt-get update && apt-get install -y ca-certificates
WORKDIR /usr/local/bin/app
COPY --from=builder /build/app .
CMD ["./app"]