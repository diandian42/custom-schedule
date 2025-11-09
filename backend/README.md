# Custom Schedule Backend

自定义计划表应用后端服务（Go语言）

## 技术栈

- Go 1.21+
- Gin Web框架
- GORM ORM框架
- JWT认证

## 项目结构

```
backend/
├── main.go              # 入口文件
├── go.mod              # 依赖管理
├── config/             # 配置文件
├── models/             # 数据模型
├── controllers/        # 控制器
├── services/           # 业务逻辑层
├── middleware/         # 中间件
├── utils/              # 工具函数
└── database/           # 数据库相关
```

## 运行

```bash
# 安装依赖
go mod download

# 运行服务
go run main.go

# 或编译后运行
go build -o app
./app
```

## 开发

服务默认运行在 `http://localhost:8080`

健康检查接口：`GET /health`

