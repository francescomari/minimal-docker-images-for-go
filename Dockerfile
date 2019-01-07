FROM golang AS build
WORKDIR /build
COPY app.go .
RUN CGO_ENABLED=0 go build -o app

FROM alpine
RUN apk add ca-certificates
COPY --from=build /build/app /app
ENTRYPOINT [ "/app" ]