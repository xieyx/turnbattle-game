# 实时对战功能开发计划

## 功能概述
实现实时对战功能，包括WebSocket实时通信、玩家匹配系统和实时战斗界面。

## 技术实现

### WebSocket实时通信
- 使用gorilla/websocket库实现WebSocket服务器
- 支持玩家连接认证和消息广播
- 实现心跳检测和连接管理

### 玩家匹配系统
- 实现匹配队列管理
- 支持随机匹配和好友对战
- 匹配成功后创建战斗房间

### 实时战斗界面
- 使用React实现动态战斗界面
- 实时显示玩家状态和战斗进度
- 支持玩家操作输入和反馈

## API设计

### WebSocket端点
```
GET /ws/battle/{battleId}
```

### REST API端点
```
POST /api/v1/matches - 创建匹配请求
GET /api/v1/matches/{matchId} - 获取匹配状态
POST /api/v1/matches/{matchId}/cancel - 取消匹配
```

## 数据模型

### Match (匹配)
- id: 匹配ID
- player1_id: 玩家1 ID
- player2_id: 玩家2 ID (可为空，表示等待匹配)
- status: 匹配状态 (waiting, matched, cancelled)
- created_at: 创建时间
- updated_at: 更新时间

## 开发任务分解

### 后端任务
1. 实现WebSocket连接管理
2. 实现匹配队列服务
3. 实现匹配API端点
4. 实现战斗房间管理
5. 实现实时消息广播

### 前端任务
1. 实现WebSocket客户端连接
2. 实现匹配界面
3. 实现实时战斗界面
4. 实现玩家操作输入
5. 实现战斗状态显示

## 测试计划
1. WebSocket连接测试
2. 匹配功能测试
3. 实时战斗测试
4. 性能压力测试
5. 兼容性测试

## 部署考虑
1. WebSocket连接数限制
2. 负载均衡配置
3. 连接超时设置
4. 错误处理和重连机制
