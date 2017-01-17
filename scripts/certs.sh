#!/bin/bash

openssl genrsa -out ../utils/certs/server.key 2048
openssl req -new -x509 -sha256 -key ../utils/certs/server.key -out ../utils/certs/server.pem -days 3650
