# 开发环境搭建指南

## 环境要求
- Go 1.19+
- Node.js 18+
- PostgreSQL 14+
- Docker (可选，用于容器化部署)

## 服务端环境搭建

### 安装Go依赖
```bash
cd server
go mod tidy
```

### 数据库配置
1. 安装PostgreSQL
2. 创建数据库:
   ```sql
   CREATE DATABASE turnbattle;
   CREATE USER turnbattle_user WITH PASSWORD 'password';
   GRANT ALL PRIVILEGES ON DATABASE turnbattle TO turnbattle_user;
   ```

3. 运行数据库迁移:
   ```bash
   psql -U turnbattle_user -d turnbattle -f migrations/001_create_users_table.sql
   ```

### 运行服务端
```bash
cd server
go run main.go
```

## 客户端环境搭建

### 安装Node.js依赖
```bash
cd client
npm install
```

### 运行客户端
```bash
cd client
npm start
```

## 开发工具推荐
- VSCode with Go and TypeScript extensions
- Postman for API testing
- pgAdmin for database management

## 项目结构说明

### 服务端 (Go)
- `main.go`: 服务端入口文件
- `models/`: 数据模型定义
- `services/`: 业务逻辑实现
- `middleware/`: 中间件
- `utils/`: 工具函数
- `migrations/`: 数据库迁移脚本

### 客户端 (React)
- `src/components/`: 可复用的UI组件
- `src/pages/`: 页面组件
- `src/services/`: API服务层
- `src/utils/`: 客户端工具函数
