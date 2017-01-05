build:
	go build -o out/multisearch main/main.go
	cp config.json out/config.json
run:
	cd out && ./multisearch