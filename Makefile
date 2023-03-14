.PHONY: build clean deploy create_env

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/cmd cmd/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

make create_env:
	sh ./create_env_file.sh