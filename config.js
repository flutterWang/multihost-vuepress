let fs = require('fs')
let file = "./docs/.vuepress/config.json"
let config = JSON.parse(fs.readFileSync(file));

let vuePressConfig = config.config

module.exports = {
  title: vuePressConfig.title,
  description: vuePressConfig.description,
  head: vuePressConfig.head,
  markdown: vuePressConfig.markdown,
  themeConfig: vuePressConfig.themeConfig
}