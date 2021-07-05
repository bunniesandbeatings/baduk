PHONY: clean
clean:
	rm baduk || true

PHONY: build
build: clean
	go build github.com/bunniesandbeatings/baduk

PHONY: install
install:
	go install github.com/bunniesandbeatings/baduk