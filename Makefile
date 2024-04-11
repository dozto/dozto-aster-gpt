build:
	go build -o ./out/gpt -tags=viper_bind_struct ./main.go
dev:
	GO_ENV=dev air
tidy:
	go mod tidy -v
