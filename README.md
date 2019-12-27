#README
仿写LyricTian作者的gin-admin项目

## Trouble shooting
- 使用wire的时候
    1. 当通过wire生成wire_gen文件时，需要将go build main.go wire_gen.go 才能启动
    2. 使用wire的时候项目内不能引入路径含有internal的文件夹