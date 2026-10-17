# syntax=docker/dockerfile:1

FROM golang:1.24-alpine AS build
ARG VERSION=dev
WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath \
	-ldflags="-s -w -X main.version=${VERSION}" \
	-o /out/hello .

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/hello /hello
ENTRYPOINT ["/hello"]
