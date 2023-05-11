package user

// import (
// 	"database/sql"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// )

// func Login(c gin.Context) {
// 	var inputUser User
// 	inputUser.Username = c.Query("username")
// 	inputUser.Password = c.Query("password")

// 	var count int //给后面的scan函数提供参数，否则会报错。此变量没有其他用处

// 	//连接mysql  "用户名:密码@[连接方式](主机名:端口号)/数据库名"
// 	db, _ := sql.Open("mysql", "root:6uqeltKH_kqwer@tcp(192.168.153.148:3306)/MyDB_one")
// 	defer db.Close()
// 	err := db.Ping()
// 	if err != nil {
// 		return
// 	}

// 	// 查询用户名和密码是否存在
// 	result := db.QueryRow("select count(*) from user_table where name= ? and password = ?;", inputUser.Username, inputUser.Password)
// 	err = result.Scan(&count)
// 	if err != nil {
// 		if err == sql.ErrNoRows {

// 			return
// 		} else {

// 			return
// 		}
// 	}

// 	//查询用户权限，存入token
// 	result = db.QueryRow("select rbac from user_table where name= ? ;", inputUser.Username)
// 	err = result.Scan(&inputUser.Permission)
// 	if err != nil {

// 		return
// 	}
// 	expiresAt := time.Now().Add(time.Minute * 30).Unix()
// 	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 	// 	"username":   inputUser.Username,
// 	// 	"expire":     expiresAt,
// 	// 	"permission": inputUser.Permission,
// 	// })

// 	//tokenString, err := token.SignedString(jwtKey)
// 	if err != nil {

// 		return
// 	}

// }
