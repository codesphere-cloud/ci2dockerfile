build:
	@go build -o yml2docker . && ./yml2docker -b alpine:latest -e MINIO_BROWSER_REDIRECT_URL=http://localhost:8089/console

run-build:
	@cd ./export && docker compose up --build

run:
	@cd ./export && docker compose up