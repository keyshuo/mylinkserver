package app

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"k8s.io/klog"
)

type MySql struct {
	mysql *sql.DB
	stmt  *sql.Stmt
}

//新建连接，从连接池拿一个
func NewMySql(msg string) (*MySql, string) {
	mysqlTemp, err := sql.Open("mysql", "root:@Wx614481987@tcp(1.15.76.132:3306)/androidDatabase")
	if err != nil {
		klog.Error(err)
		return nil, "database connection failed"
	}
	err = mysqlTemp.Ping()
	if err != nil {
		klog.Error(err)
		return nil, "database connection failed"
	}
	stmtTemp, err := mysqlTemp.Prepare(msg)
	if err != nil {
		klog.Error(err)
		return nil, "message exist problem"
	}
	return &MySql{
		mysql: mysqlTemp,
		stmt:  stmtTemp,
	}, ""
}

//用sql语句查询，返回一个结果字符串列表
func (sql *MySql) Search(args ...interface{}) ([]string, string) {
	var result []string
	var temp string
	rows, err := sql.stmt.Query(args...)
	if err != nil {
		return nil, "server search error"
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&temp)
		fmt.Println(temp)
		result = append(result, temp)
	}
	fmt.Println("search")
	return result, ""
}

func (sql *MySql) SearchRows(obj interface{}, args ...interface{}) ([]interface{}, string) {
	rows, err := sql.stmt.Query(args...)
	if err != nil {
		return nil, "server search error"
	}
	defer rows.Close()
	s := reflect.ValueOf(obj).Elem()
	t := s.Type()
	fields := make(map[string]int)
	for i := 0; i < s.NumField(); i++ {
		fields[t.Field(i).Name] = i
	}

	// 遍历查询结果
	results := make([]interface{}, 0)
	for rows.Next() {
		// 创建结构体实例
		r := reflect.New(t).Elem()

		// 将查询结果映射到结构体
		values := make([]interface{}, s.NumField())
		for i := 0; i < s.NumField(); i++ {
			field := t.Field(i)
			values[i] = r.FieldByName(field.Name).Addr().Interface()
		}
		if err := rows.Scan(values...); err != nil {
			klog.Error(err)
			return nil, "server error"
		}
		for i := 0; i < s.NumField(); i++ {
			field := t.Field(i)
			if field.Name == "Time" {
				if values[i] != nil {
					value := values[i].(*string)
					date, err := time.ParseInLocation("2006-01-02 15:04:05", *value, time.Local)
					if err != nil {
						klog.Error(err)
						return nil, "server error"
					}
					r.FieldByName(field.Name).Set(reflect.ValueOf(date.Format("2006-01-02 15:04:05")))
				}
			}
		}
		results = append(results, r.Addr().Interface())
	}

	return results, ""
}

//执行UPDATE、INSERT、DELETE操作
func (sql *MySql) Exec(args ...interface{}) string {
	_, err := sql.stmt.Exec(args...)
	if err != nil {
		klog.Error(err)
		return "operation error"
	}
	return ""
}

//关闭数据库连接
func (sql *MySql) Close() {
	err := sql.mysql.Close()
	if err != nil {
		klog.Error(err)
		return
	}
	err = sql.stmt.Close()
	if err != nil {
		klog.Error(err)
		return
	}
}

func (newSql *MySql) UpdateMysql(msg string) string {
	mysqlTemp, err := sql.Open("mysql", "root:@Wx614481987@tcp(1.15.76.132:3306)/androidDatabase")
	if err != nil {
		klog.Error(err)
		return "database connection failed"
	}
	err = mysqlTemp.Ping()
	if err != nil {
		klog.Error(err)
		return "database connection failed"
	}
	stmtTemp, err := mysqlTemp.Prepare(msg)
	if err != nil {
		klog.Error(err)
		return "message exist problem"
	}
	newSql.mysql = mysqlTemp
	newSql.stmt = stmtTemp
	return ""
}
