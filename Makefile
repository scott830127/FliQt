APP 			= fliqt
SERVER_BIN  	= ./cmd/${APP}/${APP}

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

docker-compose:
	@bash ./deploy/manifest/docker-compose/apply_ip.sh
	docker-compose up -d

dao_gen:
	go run ./tool/gorm-gen/main.go

# 編譯 linux 版用於docker-compose
docker:
	GOOS=linux GOARCH=amd64 \
	go build -o ./deploy/manifest/docker-compose/local/app/build/${APP} ./cmd/${APP}/main.go && \
	cp ./deploy/config/config.base.toml deploy/manifest/docker-compose/local/app/build/ && \
	cp -R ./deploy/file deploy/manifest/docker-compose/local/app/build && \
	cp -R ./deploy/lib deploy/manifest/docker-compose/local/app/build