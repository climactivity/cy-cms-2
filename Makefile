.PHONY: serve

.DEFAULT_GOAL := serve


all: dependencies frontend serve

dependencies: 
	npm install

frontend:
	npm run build 

serve: 
	go run cms/main.go serve

collections: 
	go run $(MAIN) migrate collections -y

MAIN = cms/main.go