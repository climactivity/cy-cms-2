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
	go run $(MAIN) migrate collections -y

run-container:
	docker-compose up -d

MAIN = cms/main.go