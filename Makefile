build:
	go build -o bin/goinode main.go

run:
	go run main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/goinode-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/goinode-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/goinode-freebsd-386 main.go
