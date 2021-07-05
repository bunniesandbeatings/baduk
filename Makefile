PHONY: clean
clean:
	rm baduk || true

PHONY: build
build: clean
	go build github.com/bunniesandbeatings/baduk