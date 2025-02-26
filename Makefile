build:
	@go build -o yml2docker . && ./yml2docker -b alpine:latest

run-build:
	@cd ./export && docker compose up --build

run:
	@cd ./export && docker compose up