## build
FROM golang:1.19-alpine AS build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# 将源代码复制到映像中
COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd

#EXPOSE 8080
#CMD ["app"]


## deploy
# FROM gcr.io/distroless/base-debian11
FROM alpine:3.17

WORKDIR /
COPY --from=build /usr/local/bin/app /usr/local/bin/app

EXPOSE 8080

#USER nonroot:nonroot
#RUN chmod -R 777 /usr/local/bin/app

ENTRYPOINT ["/usr/local/bin/app"]

