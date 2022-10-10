.PHONY: serve

.DEFAULT_GOAL := serve

serve: 
	go run cms/main.go serve

$(MAIN): main.go