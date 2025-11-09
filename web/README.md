# Custom Schedule Web

自定义计划表应用Web端（Vue3）

## 技术栈

- Vue 3
- Vite
- Vue Router
- Pinia
- Element Plus
- Axios

## 项目结构

```
web/
├── src/
│   ├── assets/        # 静态资源
│   ├── components/    # 公共组件
│   ├── views/         # 页面组件
│   ├── router/        # 路由配置
│   ├── stores/        # Pinia状态管理
│   ├── api/           # API接口
│   ├── utils/          # 工具函数
│   ├── App.vue        # 根组件
│   └── main.js        # 入口文件
├── index.html         # HTML模板
└── vite.config.js     # Vite配置
```

## 安装依赖

```bash
npm install
```

## 运行

```bash
npm run dev
```

应用将在 `http://localhost:3000` 运行

## 构建

```bash
npm run build
```

