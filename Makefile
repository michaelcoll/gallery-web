build: build-web build-go

build-go:
	go build -v .

build-web:
	cd internal/web \
	&& pnpm run build

gen: protoc

protoc:
	protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		proto/gallery.proto

clean:
	rm proto/*.pb.go

run:
	go run . index -f ~/Images/Photos

prepare:
	cd internal/web \
	&& corepack enable && corepack prepare \
	&& pnpm i