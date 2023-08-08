# 1. 背景

git commit message 是一件比较烦人的事情, 这个工具意图使用 AI 的方式根据代码生成 commit message

# 2. 项目介绍

语言: python
SDK: [dashcope](https://help.aliyun.com/zh/dashscope/developer-reference/)

dashcope 是一个阿里云灵积模型服务的一个 SDK, 使用这个 SDK 可以调用如`通义千问`等大模型, 本项目将采用该 SDK 来调用`通义千问`来生成 commit message. 而输入信息是 git diff 即代码变更结果.

# 3. 测试结果

测试结果比较失败, 生成结果基本不可用, 个人感觉通义千问(qwen)压根无法根据 git diff 的内容生产 git commit message, 换句话说就是没有训练过 git diff 的结果.

对比测试 ChatGPT 和 Claude 同样的结果提交参数效果如下:

```
qwen < chatGPT < claude
```
