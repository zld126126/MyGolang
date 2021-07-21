package main

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/labstack/gommon/log"
)

func main() {
	session, err := getDBSession()
	chkErr(err)

	// todo 第一次执行
	err = createTable(session)
	chkErr(err)

	err = insert(session)
	chkErr(err)

	err = find(session)
	chkErr(err)

	err = dropTable(session)
	chkErr(err)
}

func chkErr(err error) {
	if err == nil {
		return
	}
	log.Panic("db init error", err)
	return
}

// 拿到db
func getDBSession() (*gocql.Session, error) {
	cluster := gocql.NewCluster("127.0.0.1:9042")
	cluster.Keyspace = "dong_tech"
	cluster.Consistency = gocql.Consistency(1)
	cluster.NumConns = 3
	var err error
	dbSession, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return dbSession, nil
}

// 创建表
func createTable(session *gocql.Session) error {
	query := fmt.Sprintf(`CREATE TABLE user(id int PRIMARY KEY, user_name varchar);`)
	return session.Query(query).Exec()
}

// 删除表
func dropTable(session *gocql.Session) error {
	query := fmt.Sprintf(`drop table user;`)
	return session.Query(query).Exec()
}

// 插入数据
func insert(session *gocql.Session) error {
	query := fmt.Sprintf(`INSERT INTO user (id,user_name) VALUES (1,'zhangsan')`)
	return session.Query(query).Exec()
}

// 查询数据
func find(session *gocql.Session) error {
	query := fmt.Sprintf("SELECT * from user;")
	iter := session.Query(query).Iter()
	defer func() {
		if iter != nil {
			iter.Close()
		}
	}()
	var id int
	var name string
	for iter.Scan(&id, &name) {
		fmt.Println(id, name)
	}
	return nil
}

// 批量执行数据
func batchInsert(session *gocql.Session) error {
	query := fmt.Sprintf(`BEGIN BATCH
            UPDATE user SET user_name = 'asdqw' where id = %d;
            INSERT INTO user (id,user_name) VALUES (2,'zhangsan');
            APPLY BATCH;`, 1)
	return session.Query(query).Exec()
}
