build:
	go build -o ./out/gpt
dev:
	GO_ENV=dev air
tidy:
	go mod tidy -v
