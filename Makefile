.PHONY: all
all: presto-parse

.PHONY: presto-parse
presto-parse:
	go build -o presto-parse cmd/presto-parse/main.go

.PHONY: clean
clean:
	rm presto-parse
