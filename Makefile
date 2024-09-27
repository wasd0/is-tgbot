BIN_PATH=bin/is-tgbot/main.out

all: test build

build: 
	@go build -o ${BIN_PATH} cmd/is-tgbot/main.go

test:
	@go test -v ./...
	
run:
	@chmod +x scripts/run.sh && scripts/run.sh ${BIN_PATH}

clean:
	@chmod +x scripts/clean.sh && scripts/clean.sh ${BIN_PATH}
