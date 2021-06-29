FROM alpine

RUN mkdir -p /usr/app/vdart-service/.
WORKDIR /usr/app/vdart-service/.

COPY vdart-service .
COPY config.json ./config.json

RUN apk --no-cache add ca-certificates

CMD ["./vdart-service"]