.PHONY: dep
dep:
	@dep ensure -v -vendor-only
run:
	@go run app.go