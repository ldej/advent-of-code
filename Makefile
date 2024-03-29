# Usage `make year=2023`

setup:
	@mkdir -p $(year)
	@for i in $(foreach v, $(shell seq 1 25),$(shell printf '%02d' $(v))); do \
  		cp -r ./template $(year)/day$${i};\
	done

test:
	go test -count=1 -v ./tools/...

bench:
	(cd tools && go test -bench=.)

# Use 'make format DIR=2020' to limit the path
format:
	@(if ! [ -x "$$(command -v goimports-reviser)" ]; then \
		echo "installing github.com/incu6us/goimports-reviser/v2"; \
		go install github.com/incu6us/goimports-reviser/v2; \
	fi)

	@(for f in $$(find $(DIR). -type f -name *.go); do \
		goimports-reviser -rm-unused -file-path $$f; \
	done)
