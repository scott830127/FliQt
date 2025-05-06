APP 			= fliqt
SERVER_BIN  	= ./cmd/${APP}/${APP}
IMAGE_NAME      = ${APP}:latest

test:
	go test ./...

start:
	go run ./cmd/${APP}/main.go -b ./deploy/config/config.local.toml

wire:
	cd ./internals/app && \
	wire gen ./
	@echo "Generated wire_gen.go!"

build:
	go build -o $(SERVER_BIN) ./cmd/${APP}

docker:
	docker build -t ${IMAGE_NAME} .