# Custom Schedule Mobile

自定义计划表应用移动端（UniApp + Vue3）

## 技术栈

- UniApp
- Vue 3
- Vite
- Vuex（可选）

## 项目结构

```
mobile/
├── pages/              # 页面
│   ├── index/          # 首页
│   ├── month/          # 月计划
│   └── week/           # 周计划
├── components/         # 组件
├── static/             # 静态资源
├── store/              # Vuex状态管理
├── utils/              # 工具函数
├── api/                # API接口
├── App.vue             # 应用入口
├── main.js             # 入口文件
├── manifest.json       # 应用配置
└── pages.json          # 页面配置
```

## 安装依赖

```bash
npm install
```

## 运行

```bash
# H5端
npm run dev:h5

# 微信小程序
npm run dev:mp-weixin

# App端
npm run dev:app
```

## 构建

```bash
# H5端
npm run build:h5

# 微信小程序
npm run build:mp-weixin

# App端
npm run build:app
```

