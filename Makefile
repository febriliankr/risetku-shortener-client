deploy:
	@echo " >> Building binary (main)"
	@GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main cmd/rest/main.go
	@immortalctl halt risetku-shortener-client
	@immortal -c risetku-shortener-client.yml
