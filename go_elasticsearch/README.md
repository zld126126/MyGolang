# go_elasticsearch
go_elasticsearch

1.1、拉取镜像
docker pull elasticsearch:7.0.0

1.2 启动镜像
参考文档:https://www.jianshu.com/p/3c90f144775f
docker run --name=test_es -d -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -e "ES_JAVA_OPTS=-Xms256m -Xmx256m" -v  /Users/elasticsearch/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml docker.io/elasticsearch:7.0.0 

1.3 用法参考
https://www.bootwiki.com/elasticsearch/elasticsearch-getting-start.html
