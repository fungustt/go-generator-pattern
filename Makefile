test:
	go clean -testcache && go test -v -race ./tests
