.PHONY: serve

.DEFAULT_GOAL := serve

serve: 
	go run pb_src/main.go serve

$(MAIN): main.go