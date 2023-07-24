FROM golang:alpine as builder

WORKDIR /builder
COPY . .
RUN go mod download
RUN go build -o app
RUN ls -lh && chmod +x ./app

FROM code.hyxc1.com/open/alpine:3.16.0 as runner
LABEL org.opencontainers.image.authors="lxh@cxh.cn"

# 定义一下版本号
ARG APP_VER
ENV APP_VER=${APP_VER}

WORKDIR /app
COPY --from=builder /builder/app ./app
CMD ./app