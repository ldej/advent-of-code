test:
	go test -v ./tools/...

bench:
	(cd tools && go test -bench=.)