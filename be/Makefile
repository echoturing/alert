REPO_NAME="github.com/echoturing/alert"


.PHONY: fmt
fmt:
	@find ./ -name "*.go" | grep -v "/ent/" | xargs goimports -w -l --local $(REPO_NAME) --private "mockprivate"
	@find ./ -name "*.go" | grep  "/ent/schema" | xargs goimports -w -l --local $(REPO_NAME) --private "mockprivate"


.PHONY: run
run:
	@go run ./cmd/main.go


.PHONY: migrate
migrate:
	@go run ./cmd/main.go  --migrate=true

.PHONY: build
build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o .build/alert-be-linux cmd/main.go