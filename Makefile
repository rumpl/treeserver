all: cli

cli:
	@go build .

image:
	@docker build -t rumpl/tree .

.PHONY: cli image
