# 雷神加速器自动暂停工具

这是一个用Go语言编写的雷神加速器自动暂停工具，可以通过API自动暂停雷神加速器服务。

## 功能特性

- 🌀 自动暂停雷神加速器
- 🔄 支持重复操作检测（错误码400803）
- 🎯 简洁的命令行输出
- 📦 模块化代码结构
- ⚡ 高性能Go语言实现

## 项目结构

```
leishen-auto/
├── api/           # API客户端包
│   └── client.go  # 雷神API客户端实现
├── config/        # 配置包
│   └── config.go  # 配置加载和管理
├── main.go        # 主程序入口
├── go.mod         # Go模块文件
└── README.md      # 项目说明文档
```

## 环境要求

- Go 1.19 或更高版本
- 有效的雷神加速器账户令牌

## 安装和使用

### 1. 克隆项目

```bash
git clone <repository-url>
cd leishen-auto
```

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 设置环境变量

设置你的雷神加速器账户令牌：

```bash
# Windows (PowerShell)
$env:TOKEN="your_account_token_here"

# Windows (CMD)
set TOKEN=your_account_token_here

# Linux/macOS
export TOKEN="your_account_token_here"
```

### 4. 运行程序

```bash
go run main.go
```

或者编译后运行：

```bash
go build -o leishen-auto.exe
./leishen-auto.exe
```


## API说明


### 错误码说明

- `0`: 操作成功
- `400803`: 账号已经停止加速，请不要重复操作
- 其他错误码: 请参考雷神加速器官方API文档

## 开发

### 代码结构说明

- `api/client.go`: 包含雷神API客户端的实现，负责HTTP请求和响应处理
- `config/config.go`: 配置管理，从环境变量加载配置
- `main.go`: 主程序入口，协调各个模块完成暂停操作

### 添加新功能

如果需要添加其他雷神加速器API功能，可以在`api/client.go`中添加新的方法。

## 许可证

本项目采用MIT许可证，详情请参见LICENSE文件。