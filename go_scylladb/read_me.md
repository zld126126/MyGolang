docker命令:
1.下载docker镜像
docker pull scylladb/scylla
2.本地运行
docker run -p 9042:9042 -it --name some-scylla -d scylladb/scylla

创建数据库:
1.进入数据库:
docker exec -it some-scylla cqlsh
2.新建一个数据库
CREATE KEYSPACE IF NOT EXISTS dong_tech WITH REPLICATION = {'class': 'SimpleStrategy','replication_factor':1};
describe keyspaces;
use dong_tech;

运行本地:
main方法
可以反复执行

参考资料:
https://blog.csdn.net/niyuelin1990/article/details/79624530
