REPO_NAME="github.com/echoturing/alert"


.PHONY: fmt
fmt:
	@find ./ -name "*.go" | xargs goimports -w -l --local $(REPO_NAME) --private "mockprivate"


.PHONY: run
run:
	@go run ./cmd/main.go