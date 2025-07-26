# TurnBattle - 回合制战斗游戏

## 项目概述
TurnBattle是一个基于Go和React开发的回合制战斗模拟游戏。玩家可以在游戏中体验策略性的回合制战斗，选择不同的角色和技能进行对战。

## 项目进展
- ✅ **v0.1.0**: 用户认证系统 (2025年7月26日发布)
  - 用户注册、登录和JWT Token认证
  - 数据库迁移脚本
  - API文档和开发指南
  - 完整的后端架构搭建

- 🚧 **v0.2.0**: 战斗系统核心逻辑 (开发中)
  - 回合制战斗引擎
  - 角色属性系统
  - 技能效果计算逻辑

- 🔜 **v0.3.0**: 实时对战功能
  - WebSocket实时通信
  - 玩家匹配系统
  - 实时战斗界面

## 技术栈
- **服务端**: Go + WebSocket
- **客户端**: React + TypeScript
- **数据库**: PostgreSQL
- **部署**: Docker + Kubernetes

## 目录结构
```
.
├── server/          # Go服务端代码
├── client/          # React客户端代码
├── docs/            # 项目文档
├── templates/       # 模板文件
├── scripts/         # 脚本文件
└── README.md        # 项目说明
```

## 开发环境搭建
请参考[开发环境搭建指南](docs/development-setup.md)。

## 贡献指南
请阅读[贡献指南](docs/contributing.md)了解如何参与项目开发。

## 许可证
本项目采用MIT许可证，详情请见[LICENSE](LICENSE)文件。
