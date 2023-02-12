#! /bin/bash

sqlc generate
if [ $? -eq 127 ]; then
  echo sqlc is not locally installed. using dockerized sqlc
	docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc -f sqlc.yaml generate
fi
