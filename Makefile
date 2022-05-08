all: build

build:
	mkdir -p bin
	export SETTINGS=./config-files/app-configs.ini
	go build -o bin/app cmd/app/main.go