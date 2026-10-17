APP     := hello
VERSION ?= dev

.PHONY: build test lint run clean

build: ## 编译二进制到 bin/
	go build -trimpath -ldflags="-s -w -X main.version=$(VERSION)" -o bin/$(APP) .

test: ## 运行测试（带 race 检测和覆盖率）
	go test -race -cover ./...

lint: ## 检查格式与静态问题
	@test -z "$$(gofmt -l .)" || (echo "以下文件未格式化，请运行 gofmt -w .："; gofmt -l .; exit 1)
	go vet ./...

run: ## 本地运行
	go run .

clean: ## 清理构建与测试产物
	rm -rf bin dist coverage.out coverage.html
