# Start by building the application.
FROM golang:1 as build-go

ARG BUILDTIME
ARG VERSION
ARG REVISION

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/gallery-web -ldflags="-s -w -X 'github.com/michaelcoll/gallery-web/cmd.version=$VERSION'"

# Now copy it into our base image.
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build-go /go/bin/gallery-web /bin/gallery-web

EXPOSE 8080
EXPOSE 9000

CMD ["gallery-web", "serve"]
