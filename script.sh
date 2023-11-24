#!/bin/bash

docker run -d --name healthcheck \
-p 127.0.0.1:5555:4444 \
-e DB_HOST=postgresql \
-e DB_PORT=5432 \
-e DB_USER=postgres \
-e DB_PASS=root \
-e DB_NAME=postgres \
--network=postgresql \
achimonchi/healthcheck:2.0