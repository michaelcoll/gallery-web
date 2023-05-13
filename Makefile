build: build-web build-go

prepare:
	cd internal/web \
	&& corepack enable && corepack prepare \
	&& pnpm i

dep-upgrade: dep-upgrade-go dep-upgrade-node

dep-upgrade-prepare:
	git checkout main && \
	git pull && \
	git checkout -b chore/dep-upgrade

dep-upgrade-go:
	go get -u
	go mod tidy -go -e

dep-upgrade-node:
	cd internal/web \
  && pnpm update --latest

dep-upgrade-post:
	git add . && \
	git commit -m "chore(dep): Upgrade dependency versions" && \
	git push origin main && \
	gh pr create --fill --base main --label dependencies --label go --label javascript


build-go:
	go build -v -ldflags="-s -w -X 'github.com/michaelcoll/gallery-web/cmd.version=v0.0.0'" .

build-web:
	cd internal/web \
	&& pnpm run build

build-docker:
	docker build . -t web --pull --build-arg VERSION=v0.0.1

.PHONY: test
test:
	go test -v ./...

run:
	go run . serve

run-vue:
	cd internal/web \
  && pnpm run dev

vue-lint:
	cd internal/web \
  && pnpm run lint

run-docker:
	docker run -ti --rm -p 8080:8080 web:latest
