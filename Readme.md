# yml2docker

## Build tool
Run `make build` in a terminal of this folder to create the go build.

## Use tool
To export the example run `make example-export-single` and `make example-export-multi` for an old and new ci.yml. To run the newly created example docker compose file run `make run` (only uses the multi/new export).

To use your own command you can use `./yml2docker -b ... -i ... -o ... -e ...`.

Available parameters are:
- `-b`: Base image for the dockerfile. **(Required)**
- `-i`: Input path for the **ci.yml** file. Default is `./ci.yml`.
- `-o`: Output path of the folder including docker compose and services. Default is `./export`.
- `-e`: Env vars to put into docker compose services. Multiple can be added via multiple `-e` arguments.