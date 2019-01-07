# Minimal Docker images for Go

This repository shows how to create minimal Docker images for Go applications. You need to be aware of the following points.

1. When building the application, use `CGO_ENABLED=0`. This disables cgo, which is a mechanism used by Go to call into C code. You can find more information about cgo [here|https://golang.org/cmd/cgo/].
2. You want to use the most minimal base image for the final build step. `scratch` can often be enough, but you might have problems to run your applications if you use SSL certificates, e.g. when performing an HTTPS request.
3. If SSL certificates are a problem, you need to use a base image which includes either your own CA certificates or the most common ones. In this repository I decided to use `alpine` as a base image and to install `ca-certificates` via `apk`.

# Usage

Build the image:

    docker build -t app .

Run the image:

    docker run app

If everything works as expected, you will see a dump of the Google home page.

# License

This is free and unencumbered software released into the public domain.