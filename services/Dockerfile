FROM alpine:latest

RUN apk add -U --no-cache ca-certificates
WORKDIR /home/saisatch
ADD build/builderTools .
ADD build/config* ./
ADD grpcwebproxy .

ENV RUN_ENV=prod

CMD ./builderTools; ./grpcwebproxy --run_tls_server=false --allow_all_origins --backend_addr=localhost:50051
