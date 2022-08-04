#!/usr/bin/env bash

ganache_port=8545

ganache_running() {
  nc -z localhost "$ganache_port"
}

start_ganache() {
  ./node_modules/.bin/iganache-cli -h 0.0.0.0 -l 2000000000000000000 -p $ganache_port -e 10000000000000000000 -m gravity top burden flip student usage spell purchase hundred improve check genre &
  ganache_pid=$!
  echo "ganache-cli started with pid $ganache_pid"
  echo $ganache_pid > ganache.pid
}

if ganache_running; then
  echo "Using existing ganache instance at port $ganache_port"
else
  echo "Starting our own ganache instance at port $ganache_port"
  start_ganache
fi
ls
