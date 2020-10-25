let config = require('./config.js')

let isExistMoreNav = false
let moreNavIndex = 0
for (let i = 0; i < config.themeConfig.nav.length; i++) {
  if (config.themeConfig.nav[i].items != undefined) {
    isExistMoreNav = true
    moreNavIndex = i
  }
}

if (isExistMoreNav) {
  config.themeConfig.nav[moreNavIndex].items.push({
    text: "GOCN",
    link: "https://gocn.vip/"
  })
} else {
  config.themeConfig.nav.push({
    text: "了解更多",
    items: [{
      text: "GOCN",
      link: "https://gocn.vip/"
    }],
  })
}

config.head.push([
  "script",
  {
    type: "text/javascript"
  },
  `var _hmt = _hmt || [];
    (function() {
        var hm = document.createElement("script");
        hm.src = "https://hm.baidu.com/hm.js?c77f15742ac7b6883fb18421ee33a702";
        var s = document.getElementsByTagName("script")[0];
        s.parentNode.insertBefore(hm, s);
    })();
  `
])

var formattedStr = JSON.stringify(config, null, 2);

let fs = require('fs')
let contents = ['module.exports =', formattedStr]

function writeFile(filePath) {
  fs.writeFileSync(filePath,'');
  for (let i = 0; i < contents.length; i++) {
    fs.appendFileSync(filePath, contents[i]);
  }
}

writeFile('config.js')