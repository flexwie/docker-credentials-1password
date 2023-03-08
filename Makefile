build:
	GOOS=darwin GOARCH=arm64 go build -o ./out/op-docker ./cmd/