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
