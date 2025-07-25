# TurnBattle Server

## 目录结构
```
.
├── main.go              # 服务端入口文件
├── go.mod              # Go模块定义
├── models/             # 数据模型
├── services/           # 业务逻辑服务
├── middleware/         # 中间件
├── utils/              # 工具函数
├── migrations/         # 数据库迁移脚本
└── README.md           # 服务端说明文档
```

## 数据库设置

1. 安装PostgreSQL
2. 创建数据库和用户:
   ```sql
   CREATE DATABASE turnbattle;
   CREATE USER turnbattle_user WITH PASSWORD 'password';
   GRANT ALL PRIVILEGES ON DATABASE turnbattle TO turnbattle_user;
   ```

3. 运行数据库迁移:
   ```bash
   psql -U turnbattle_user -d turnbattle -f migrations/001_create_users_table.sql
   ```

## 运行服务端

```bash
cd server
go run main.go
```

服务将在 `localhost:8080` 上运行。

## API端点

### 公共端点
- `GET /` - API根路径
- `POST /api/v1/users/register` - 用户注册
- `POST /api/v1/users/login` - 用户登录
- `GET /ws` - WebSocket连接

### 需要认证的端点
- `GET /api/v1/users/profile` - 获取用户信息

## 环境变量

在生产环境中，应该设置以下环境变量:
- `DB_HOST` - 数据库主机
- `DB_PORT` - 数据库端口
- `DB_USER` - 数据库用户
- `DB_PASSWORD` - 数据库密码
- `DB_NAME` - 数据库名称
- `JWT_SECRET` - JWT密钥
