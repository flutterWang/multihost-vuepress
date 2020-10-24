let fs = require('fs')
let file = "./docs/.vuepress/config.json"
let config = JSON.parse(fs.readFileSync(file));

let vuePressConfig = config.config

vuePressConfig.themeConfig.nav.push(
  {
    text: "了解更多",
    items: [
      { text: "GOCN", link: "https://gocn.vip/" }
    ],
  }
)

module.exports = {
  title: vuePressConfig.title,
  description: vuePressConfig.description,
  head: vuePressConfig.head,
  markdown: vuePressConfig.markdown,
  themeConfig: vuePressConfig.themeConfig
}