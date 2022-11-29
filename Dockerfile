# Start by building the application.
FROM node:18 as build-node

WORKDIR /node/src
COPY . .

RUN cd internal/web \
    	&& corepack enable && corepack prepare \
    	&& pnpm i \
      && pnpm run build

FROM golang:1.19 as build-go

ARG BUILDTIME
ARG VERSION
ARG REVISION

WORKDIR /go/src/app
COPY . .
COPY --from=build-node /node/src/internal/web/dist /go/src/app/internal/web/dist

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/gallery-web -ldflags="-s -w -X 'github.com/michaelcoll/gallery-web/cmd.version=$VERSION'"

# Now copy it into our base image.
FROM gcr.io/distroless/static-debian11:nonroot

COPY --from=build-go /go/bin/gallery-web /bin/gallery-web

EXPOSE 8080
EXPOSE 9000

CMD ["gallery-web", "serve"]
