FROM golang:1.19.4 as build_context

RUN mkdir -p /out/usr/local/bin/

WORKDIR /workspace

ADD . /workspace/

RUN go build -o api-test ./main.go

RUN mv /workspace/api-test /out/usr/local/bin/

##############
# Final image
#############

FROM alpine:3.17

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY --from=build_context /out/usr/local/bin/api-test /

WORKDIR /

EXPOSE 9090

CMD ["./api-test"]