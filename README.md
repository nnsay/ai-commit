# 1. 项目介绍

写提交信息和给变量起名一样, 对于程序员来说都是头疼的事情, 本项目使用 Google 的 Gemini 服务根据本地的 git diff 结果自动生成合适的提交信息.

# 2. 项目特点

- Gemini 比较先进多模态框架, 号称比 ChatGPT4 好
- 简单, 核心代码<100 行
- 标准化, 生成的提交信息符合[约定式提交](https://www.conventionalcommits.org/zh-hans/v1.0.0/)
- 支持 Monorepo(nx)项目结构

# 3. 使用方法

该工具是一个命令行工具, 下载代码在本地构建, 或者使用已经发布的 Release.

## 3.1 配置

需要配置`GEMINI_API_KEY`环境变量

## 3.2 执行

```bash
# 只需调用命令, 即可获取自动生成提交信息
ai-commit

# 支持增加额外参数的 git diff 参数
ai-commit --staged
```

# 4. 许可

MIT
