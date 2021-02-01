test:
	go test -count=1 -v ./tools/...

bench:
	(cd tools && go test -bench=.)

format:
	for f in $$(find . -type f -name *.go); do \
		goimports-reviser -rm-unused -file-path $$f; \
	done