package app

import (
	"database/sql"

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
		result = append(result, temp)
	}
	return result, ""
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
