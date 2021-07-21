#安装步骤
1.go get github.com/sclevine/agouti

2.go get github.com/onsi/ginkgo/ginkgo

3.go get github.com/onsi/gomega

4.brew tap homebrew/cask

5.brew cask install chromedriver
[可能需要翻墙]

6.gin
go get github.com/gin-gonic/gin

#brew 翻墙问题
export ALL_PROXY=socks5://192.168.4.144:10880

#代码：
terminal运行:

ginkgo bootstrap --agouti

ginkgo generate --agouti user_login

#Proxy
export GO111MODULE=on
export GOPROXY=https://goproxy.io

#测试步骤
1运行main.go
2运行*_suite_test.go


