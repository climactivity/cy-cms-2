.PHONY: serve

.DEFAULT_GOAL := serve

serve: 
	go run cms/main.go serve

collections: 
	go run $(MAIN) migrate collections -y

MAIN = cms/main.go