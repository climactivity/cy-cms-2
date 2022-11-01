.PHONY: serve

.DEFAULT_GOAL := all


all: dependencies frontend serve

check-compiles: dependencies frontend build-go

build-go: 
	go build $(MAIN) 

dependencies: 
	npm install

frontend:
	npm run build 

serve: 
	go run cms/main.go serve

collections: 
	go run $(MAIN) migrate collections

run-container:
	docker-compose up -d

build-container: 
	docker build . -f ./Dockerfile --tag ghcr.io/climactivity/cy-cms-2:latest
	docker push ghcr.io/climactivity/cy-cms-2:latest

MAIN = cms/main.go