work_path=$(pwd)
mkdir ../../quick_vuepress && cd ../../quick_vuepress
npm init -y
mkdir docs && cd docs && cp ${work_path}/README.md ./
mkdir .vuepress && cd .vuepress && cp ${work_path}/config.js ./
cp ${work_path}/config.json ./