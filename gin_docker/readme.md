1.编译main
sudo env GOOS=linux GOARCH=amd64 go build -o main

2.docker build
docker build -t landon:v1 .

3.1启动docker 方法1：
docker run --rm -landon_web -p 8991:8991 -p 8992:8992 landon:v1 

3.2启动docker 方法2：
docker-compose -f docker-compose.yml up -d landon_web
