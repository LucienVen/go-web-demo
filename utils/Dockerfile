## build
FROM golang:1.19-alpine AS build

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
WORKDIR /usr/src/app
COPY . .

#EXPOSE 8080
#CMD ["app"]
