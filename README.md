## 执行
1. make go
2. 在生成 quick_vuepress 里执行 npm install -D vuepress
3. 在生成项目的 package.json 中的 scripts 中加入 
```
 "docs:dev": "vuepress dev docs",
  "docs:build": "vuepress build docs"
```
4. 执行 npm run docs:dev 启动程序