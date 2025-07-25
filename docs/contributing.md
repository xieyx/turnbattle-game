# 贡献指南

感谢您对TurnBattle项目的关注！我们欢迎任何形式的贡献。

## 贡献方式

### 报告Bug
- 使用GitHub Issues报告Bug
- 请提供详细的复现步骤和环境信息

### 提交功能请求
- 使用GitHub Issues提交功能请求
- 请详细描述功能的用途和实现思路

### 代码贡献
1. Fork项目到您的GitHub账户
2. 创建功能分支 (`git checkout -b feature/your-feature-name`)
3. 提交更改 (`git commit -am 'Add some feature'`)
4. 推送到分支 (`git push origin feature/your-feature-name`)
5. 创建Pull Request

## 代码规范

### Go代码规范
- 遵循Go官方代码风格
- 使用`gofmt`格式化代码
- 编写单元测试，测试覆盖率需达到80%以上

### TypeScript/React代码规范
- 遵循Airbnb JavaScript规范
- 使用ESLint和Prettier进行代码检查和格式化
- 使用TypeScript编写类型安全的代码

## Git提交规范
- 使用清晰的提交信息
- 遵循Conventional Commits规范
- 每个提交应该只包含一个逻辑变更

## Pull Request流程
1. 确保代码通过所有测试
2. 更新相关文档
3. 添加适当的标签
4. 指定合适的审阅者

## 代码审查
- 所有PR都需要至少两名审阅者批准
- 审阅者应关注代码质量、性能和安全性
- 及时响应审阅意见

## 测试要求
- 单元测试覆盖率不低于80%
- 集成测试覆盖核心功能
- 性能测试确保响应时间符合要求
