#!/bin/bash


docker build --rm --compress --tag bodsch/jolokia_exporter:latest .

docker create -ti --name dummy bodsch/jolokia_exporter:latest bash
docker cp dummy:/root/jolokia_exporter bin/
docker rm -fv dummy

