FROM node:12 as node-builder

WORKDIR /build

COPY ./ui /build

RUN npm install . && npm run build

FROM golang:1.14 AS golang-builder

WORKDIR /build

COPY . /build

RUN go build -o hackine --ldflags "-linkmode 'external' -extldflags '-static'" .

FROM alpine:3

WORKDIR /bin

COPY --from=node-builder /build/dist /usr/share/hackine/assets
COPY --from=golang-builder /build/hackine /bin/hackine

EXPOSE 80

ENTRYPOINT [ "/bin/hackine" ]
