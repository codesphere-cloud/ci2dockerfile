# yml2docker

## Build tool
Run `make build` in a terminal of this folder.

## Use tool
Run for example `./yml2docker -b alpine:latest`.

Available parameters are:
- `-b`: Base image for the dockerfile. **(Required)**
- `-i`: Input path for the **ci.yml** file. Default is `./ci.yml`.
- `-o`: Output path of the folder including docker compose and services. Default is `./export`.