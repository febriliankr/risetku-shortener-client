build:
	@echo " >> Building binary (main)"
	@PORT=5500 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main cmd/rest/main.go
