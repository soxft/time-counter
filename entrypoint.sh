#!/bin/sh

set -x

cd /app || exit

if [ ! -f "/app/expose/install.lock" ];then
  if [ -n "$API_SERVER" ];then
    sed -i "s/http:\/\/localhost:8080\/counter/$API_SERVER/g" dist/counter.js
  fi

  mv dist /app/expose/dist
  touch /app/expose/install.lock
fi

exec /app/time-counter -c /app/config.yaml -d ./expose/dist