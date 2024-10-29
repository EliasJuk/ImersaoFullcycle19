# RUN PROJECT
go run .\cmd\videoconverter\main.go

# DOCKER

## Criar volume e criar container
$ docker volume create external-storage
$ docker compose up -d

## Acessar Container
$ docker compose exec -it go_app_dev bash


## OTHERS

docker ps
docker rm -f 04fccceb6b69

# copia arquivos para o diretorio media
cp -R /app/mediatest/media/uploads/1/* .