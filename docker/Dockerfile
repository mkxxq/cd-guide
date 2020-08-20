FROM alpine:3
RUN apk add --no-cache curl busybox-extras

EXPOSE 1323
RUN mkdir -p /app
WORKDIR /app
ADD _build/helloworld /app/helloworld

CMD ./helloworld