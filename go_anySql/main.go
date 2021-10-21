package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
	"time"
)

func main() {
	mysql()
}

var (
	MysqlUserName = "root"
	MysqlPassword = "123456"
	MysqlAddress  = "localhost"
	MysqlPort     = 3306
	MysqlDB       = "dongbao"
	Mysql         = "mysql"
)

func mysql() {
	connect := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		MysqlUserName, MysqlPassword, MysqlAddress, MysqlPort, MysqlDB)
	db, err := gorm.Open(Mysql, connect)
	if err != nil {
		log.Fatal(err)
	}
	db.DB().SetConnMaxLifetime(time.Minute * 5)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(20)
	db.LogMode(true)
	// 默认需要关闭
	defer db.Close()

	// 查询
	Query_Examples(db, "select * from users")
	// 执行
	Exec_Examples(db, "update users set name = 'dong' where id = 1")
}

type DatabaseResult struct {
	Tp    map[string]string
	Value map[string][]byte
}

func Query_Examples(db *gorm.DB, sqlStr string) {
	rows, err := db.Raw(sqlStr).Rows()
	if err != nil {
		log.Fatal(err)
	}

	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	types, err := rows.ColumnTypes()
	if err != nil {
		log.Fatal(err)
	}

	values := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for i, _ := range cols {
		scans[i] = &values[i]
	}

	i := 0
	result := make(map[int]*DatabaseResult)
	for rows.Next() {
		err = rows.Scan(scans...)
		if err != nil {
			log.Fatal(err)
		}

		tp := make(map[string]string)
		row := make(map[string][]byte)
		j := 0
		for k, v := range values {
			key := cols[k]
			//这里把[]byte根据条件转换
			row[key] = v
			tp[key] = types[j].DatabaseTypeName()
			j++
		}

		result[i] = &DatabaseResult{
			Tp:    tp,
			Value: row,
		}
		i++
	}

	for i := 0; i < len(result); i++ {
		for k, v := range result[i].Value {
			tp := result[i].Tp[k]
			switch tp {
			case "TINYINT", "SMALLINT", "INT", "INTEGER":
				s := string(v)
				if s != "" {
					i, err := strconv.Atoi(s)
					if err == nil {
						fmt.Println(i)
						// TODO 逻辑处理
					}
				}
				break

			case "DECIMAL", "FLOAT", "DOUBLE", "BIGINT":
				s := string(v)
				if s != "" {
					f, err := strconv.ParseFloat(s, 64)
					if err == nil {
						fmt.Println(f)
						// TODO 逻辑处理
					}
				}
				break
			case "VARCHAR", "CHAR", "TEXT":
				s := string(v)
				fmt.Println(s)
				break
				// TODO 逻辑处理
			default:
				fmt.Println("unknown or unSupport type in mysql")
				break
				// TODO 逻辑处理
			}

		}
	}

	for k, v := range result {
		fmt.Println("Query_Examples:第"+fmt.Sprint(k+1)+"对象"+",值:", v)
		// TODO 逻辑处理
	}
}

func Exec_Examples(db *gorm.DB, sqlStr string) {
	db = db.Exec(sqlStr)
	rowsAffected := db.RowsAffected

	err := db.Error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Exec_Examples RowsAffected:", rowsAffected)
}
