.PHONY: fmt config

fmt:
	gocmt -t "TODO" -i -d .
	gofmt -s -w .

config:
	go install github.com/cuonglm/gocmt
	go install github.com/segmentio/golines
