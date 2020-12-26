# Go parameters
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

all: test build

create_dir:
	rm -rf target/
	mkdir target/

only_build:
	$(GOBUILD) -o target/web api/main.go

build: create_dir only_build doc
	cp api/*.toml target/
	cp -R api/swagger target/

test:
	$(GOTEST) -v ./...

clean:
	rm -rf target/

run:
	target/web -conf="./target/conf.toml"

stop:
	pkill -f target/web

doc:
	qbtool swag init -d=api -g=main.go -o=api/swagger
db:
	qbtool db reverse -source="root:e23456@tcp(172.16.3.21:8306)/cherry?charset=utf8" -path="./"
	

