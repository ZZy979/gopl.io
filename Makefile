all:
	for d in $(wildcard ch*); do \
	  mkdir -p out/$$d; \
	  go build -o out/$$d ./$$d/...; \
	done

test: all
	for d in $(wildcard ch*); do \
	  go test -v ./...; \
	done

clean:
	rm -r out

.PHONY: all test clean
