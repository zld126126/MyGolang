# auto_api
auto_api 自动化后端api接口程序生成后端exe

## 具体说明文档:
    auto-api 
        |-- with golang env 需要golang环境
        |-- without golang env 不需要golang环境

## 写在前面的话:
    最近正在和前端同学一起开发,作为一个后端老鸟,自然需要给他们提供一个好用的get/post接口工具了,所以我就想到配置一个自定义接口生成工具,又简单又好用的那种
    目前有两个版本:
        1.需要golang环境的html配置
            auto-api(with golang env).exe
        2.不需要golang环境的json配置
            auto-api(without golang env).exe

## 用法很简单
    需要golang环境的版本1:
        1.1安装golang的环境:https://golang.org/dl/
        1.2运行auto-api(with golang env).exe
        1.3访问http://localhost:10086
        1.4配置并生成后端exe即可
        1.5运行生成的exe其实就是自定义api了...
    
    无需golang环境的版本2:
        2.1配置config.json
        2.2运行auto-api(without golang env).exe
        2.3运行的结果其实就是自定义api了...