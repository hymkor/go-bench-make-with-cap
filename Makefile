all:
	go fmt
	go test -bench . -benchmem
