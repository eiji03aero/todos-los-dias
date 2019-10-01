#!/bin/bash

container_name="todos-los-dias"
COMMAND=${1:-up}

function execute-docker-compose () {
  docker-compose \
    -f 'docker-compose.dev.yml' \
    $@
}

function execute-docker-sync () {
  docker-sync \
    $@ \
    -c 'docker-sync.yml'
}

function stop-docker-compose () {
  execute-docker-sync stop
  execute-docker-compose stop
}

if [ $COMMAND = 'up' ] && [ $# -le 1 ]; then
  # trap 'stop-docker-compose' SIGINT
  # execute-docker-sync start
  # execute-docker-compose up

  execute-docker-sync start
  execute-docker-compose up -d
  execute-docker-compose exec $container_name bash
  stop-docker-compose
elif [ $COMMAND = 'bash' ]; then
  execute-docker-compose exec $container_name bash

elif [ $COMMAND = 'db-init' ]; then
  docker cp ./scripts/db-init.sh cockroach:/cockroach
  docker cp ./sql/db-init.sql cockroach:/cockroach
  execute-docker-compose exec cockroach /bin/sh ./db-init.sh

else
  execute-docker-compose $@
fi
