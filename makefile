build_linux_x64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o builds/speedTester_linux_x64 .
