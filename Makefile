test:
	go test -count=1 -v ./tools/...

bench:
	(cd tools && go test -bench=.)