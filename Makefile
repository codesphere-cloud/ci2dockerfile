build:
	@go build -o yml2docker .

example-export-single:
	@./yml2docker -b alpine:latest -i ./ci.single.yml -o ./export/single

example-export-multi:
	@./yml2docker -b alpine:latest -i ./ci.multi.yml -o ./export/multi -e MINIO_BROWSER_REDIRECT_URL=http://localhost:8089/console

run:
	@cd ./export/multi && docker compose up

run-build:
	@cd ./export/multi && docker compose up --build

release:
	goreleaser release