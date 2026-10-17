# cicd_demo

一个用 Go 写的 Hello World，附带完整的 CI/CD 流水线示例。

## 本地运行

```bash
go run .            # 输出 Hello, World!
go run . Gopher     # 输出 Hello, Gopher!
go run . -version   # 输出版本号（默认 dev）
```

也可以使用 Makefile：

```bash
make run     # 运行
make test    # 测试（-race + 覆盖率）
make lint    # gofmt + go vet
make build   # 编译到 bin/hello
```

## 项目结构

```
.
├── main.go                    # 入口：Hello World + 版本号注入
├── main_test.go               # 单元测试
├── go.mod
├── Makefile                   # 本地开发命令
├── Dockerfile                 # 多阶段构建，distroless 运行镜像
├── .goreleaser.yaml           # 跨平台发布配置
└── .github/workflows/
    ├── ci.yml                 # CI：lint / test / build
    └── release.yml            # CD：打 tag 后发布二进制 + 推送镜像
```

## CI（持续集成）

`.github/workflows/ci.yml` 在推送到 `main` 或发起 PR 时触发，包含三个任务：

| Job | 内容 |
| --- | --- |
| Lint | `gofmt` 格式检查、`go vet` 静态检查 |
| Test | `go test -race`，生成覆盖率并上传 artifact |
| Build | 编译二进制、冒烟测试、上传构建产物（依赖 Lint/Test 通过） |

## CD（持续交付）

`.github/workflows/release.yml` 在推送 `v*` 形式的 tag 时触发：

1. **GoReleaser**：交叉编译 linux / darwin / windows × amd64 / arm64，生成压缩包和 checksums，创建 GitHub Release。
2. **Docker**：基于发布 tag 构建镜像并推送到 GHCR，标签为 `1.2.3`、`1.2`、`1`。

### 发布一个版本

```bash
git tag v1.0.0
git push origin v1.0.0
```

镜像地址：`ghcr.io/zhengsjcode/cicd_demo:1.0.0`

```bash
docker run --rm ghcr.io/zhengsjcode/cicd_demo:1.0.0
```

## 说明

- 版本号在构建时通过 `-ldflags "-X main.version=..."` 注入，`go run` 时显示 `dev`。
- Release 使用仓库自带的 `GITHUB_TOKEN`，无需额外配置 secrets。
- 如果 GHCR 镜像首次推送后拉取报权限错误，请在仓库 **Settings → Packages** 中确认该包的可见性/继承权限。
