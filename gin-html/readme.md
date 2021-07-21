## 普通 运行步骤
1.server目录下server.exe
2.访问localhost:9090/index

### docker 运行
3.docker访问
3.1编译main
sudo env GOOS=linux GOARCH=amd64 go build -o main
set GOOS=linux set GOARCH=amd64 go build -o main
3.2docker build
docker build -t landon:v1 .
3.3启动docker 方法1：
docker run --rm -landon_web -p 9090:9090 landon:v1 
3.4启动docker 方法2：
docker-compose -f docker-compose.yml up -d landon_web

### TODO
1.发布文章功能
2.追评功能
