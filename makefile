.SILENT:

build:
	go build -o ./bin/flaggy cmd/http-server/main.go
	GOOS=linux GOARCH=amd64 go build -o ./bin/bootstrapper cmd/http-server/main.go
dockerize: build
	docker build -f ./dist/local/Dockerfile . -t flaggy:latest

start: dockerize
	docker compose -f ./dist/local/docker-compose.yml down --remove-orphans
	docker compose -f ./dist/local/docker-compose.yml up -d

create-flag:
	curl -X DELETE http://localhost:8088/featureflag/my-feature
	curl -X POST http://localhost:8088/featureflag --header "Content-Type=application/json" --data '{"key": "my-feature", "enabled" : false }'
	curl http://localhost:8088/featureflag/my-feature

enable-flag:
	curl -v -X PUT http://localhost:8088/featureflag/my-feature/true
	curl http://localhost:8088/featureflag/my-feature

disable-flag:
	curl -X PUT http://localhost:8088/featureflag/my-feature/false
	curl http://localhost:8088/featureflag/my-feature