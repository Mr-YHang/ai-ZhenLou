### 枕楼管家 -- 基于Eino框架实现的ReAct Agent

最近《藏海传》挺火，用这个为背景来全代码开发一个ai智能体，他可以实现：

- [x] 与客人交流，帮客人推荐枕楼的菜品、点菜、介绍菜品的做法、历史文化等
- [x] 与老板交流，帮老板计算营收
- [x] 客人等待期间，可以让八公子来说评书



学习目标：

1. ollama本地部署私有化模型
2. 模型微调，让模型学习私有数据《枕楼员工手册》 todo
3. ~~Rag 构建每日菜单、每日特价菜~~  模拟内部交易系统数据入库，直接用mcp去智能调用mysql数据，后续会学习kag看看
4. 工具封装，智能调用：
   - 官方工具：如google search , 用来搜索菜品历史文化与做法；
   - 公司内部其他项目接口：如八公子的评书知识库、未来帮助用户下单结账调用交易系统等
   - mcp数据库直查
5. 上下文历史存储，并在同一个会话中的每次问答都实现上下文记忆
