# Base image is in https://registry.hub.docker.com/_/golang/
# Refer to https://blog.golang.org/docker for usage
FROM golang:1.15.5-alpine as base
WORKDIR /go/src/github.com/gfonseca/conductor
EXPOSE 3000

# Enforce to use UTF8 char code
ENV LANG en_US.UTF-8
ENV LC_ALL=C
ENV LANGUAGE en_US.UTF-8


#### TEST
FROM base as test
CMD ["make", "test"]


# #### MIGRATION
# FROM base as migration

# COPY . .

# RUN pwd && cd cmd/trinity/migration && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
#       -ldflags='-w -s -extldflags "-static"' -a \
#       -o /go/bin/migration .

# ENV USER=trinity

# RUN addgroup --gid 1001 --system $USER && adduser -u 1001 --system $USER --gid 1001

# RUN chown $USER:$USER /go/bin/migration -R

# ENTRYPOINT ["/go/bin/migration"]


#### Development
FROM base AS dev

RUN apk add --no-cache git

# Dev Dependencies
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go get -u github.com/tsenart/vegeta

# application and watches for changes
ARG FRESHER_VERSION=1.2.1

RUN wget -c https://github.com/roger-russel/fresher/releases/download/v${FRESHER_VERSION}/fresher_${FRESHER_VERSION}_Linux_x86_64.tar.gz \
      && tar xzf fresher_${FRESHER_VERSION}_Linux_x86_64.tar.gz -C /go/bin/ \
      && rm -f fresher_*tar.gz

CMD ["fresher", "-c", "./fresher.yaml"]

EXPOSE 33333


#### BUILD
FROM base as build
RUN apk add --no-cache make gcc libc-dev

COPY . .
RUN make build


### FINAL
FROM alpine AS final

COPY --from=build /build/server /server

ENTRYPOINT [ "/server" ]