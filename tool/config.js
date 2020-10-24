module.exports = {
  title: 'Jupiter',
  description: 'Framework',
  head:[
  
    [
      'meta',
      
        {
          name: 'keywords',
          content: 'Go,golang,gRPC'
        },
      
    ],
  
  ],
  markdown: {
    vuePressConfig: true,
  },
  themeConfig: {
    nav:[
      
        {
          text: '首页',
          link: '/',
        },
      
        {
          text: '文档',
          link: '/jupiter/',
        },
      
      {
        text: "了解更多",
        items: [
          { text: "GOCN", link: "https://gocn.vip/" }
        ],
      }
    ],
    sidebar:{
      
        '/jupiter/':[
          
            {
              title: '第1章 Jupiter简介',
              collapsable: false,
              children:[
                
                  '/jupiter/1.1quickstart',
                
              ]
            },
          
        ],
      
    }
  }
}