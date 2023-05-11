package app

import (
	"database/sql"

	"k8s.io/klog"
)

type MySql struct {
	mysql *sql.DB
}

//新建连接，从连接池拿一个
func NewMySql() (*MySql, error) {
	mysqlTemp, err := sql.Open("mysql", "root:@Wx614481987@tcp(1.15.76.132:3306)/lianjiemingcheng")
	return &MySql{
		mysql: mysqlTemp,
	}, err
}

//防止SQL注入
func (sql *MySql) Prepare(msg string) error {
	_, err := sql.mysql.Prepare(msg)
	return err
}

//用sql语句查询，返回一个结果字符串列表
func (sql *MySql) Search(msg string) ([]string, error) {
	var result []string
	var temp string
	rows, err := sql.mysql.Query(msg)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&temp)
		result = append(result, temp)
	}
	return result, err
}

//执行UPDATE、INSERT、DELETE操作
func (sql *MySql) Exec(msg string) error {
	_, err := sql.mysql.Exec(msg)
	return err
}

//测试数据库连接
func (sql *MySql) Test() error {
	return sql.mysql.Ping()
}

//关闭数据库连接
func (sql *MySql) Close() error {
	err := sql.mysql.Close()
	if err != nil {
		klog.Error(err)
	}
	return err
}
