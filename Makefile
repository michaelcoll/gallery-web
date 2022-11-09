build: build-web build-go

build-go:
	go build -v -ldflags="-s -w -X 'github.com/michaelcoll/gallery-web/cmd.version=v0.0.0'" .

build-web:
	cd internal/web \
	&& pnpm run build

.PHONY: test
test:
	go test -v ./...

run:
	go run . serve

prepare:
	cd internal/web \
	&& corepack enable && corepack prepare \
	&& pnpm i

dep-upgrade:
	go get -u
	go mod tidy