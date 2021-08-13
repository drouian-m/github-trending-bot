COVERAGE_DIR=coverage
MAIN=main.go
PKG_NAME=github-trending-bot

build:
	go build -ldflags "-w -s" -o $(PKG_NAME) $(MAIN)
	upx -8 $(PKG_NAME)

unused:
	go mod tidy

run: build
	./$(PKG_NAME)
