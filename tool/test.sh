work_path=$(pwd)
mkdir ../../quick_vuepress && cd ../../quick_vuepress
npm init -y
mkdir docs && cd docs && cp ${work_path}/tool/README.md ./
mkdir .vuepress && cd .vuepress && cp ${work_path}/tool/config.js ./