package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:qq757467810@tcp(localhost:3306)/jxgl")
	if err != nil {
		return nil, err
	}

	// 验证数据库连接
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
func getDataHandler(c echo.Context) error {
	// 连接数据库
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT Sname FROM student")
	if err != nil {
		return err
	}
	defer rows.Close()

	// 构造 HTML 内容
	var htmlContent string
	for rows.Next() {
		var data string
		err := rows.Scan(&data)
		if err != nil {
			return err
		}
		htmlContent += "<p>" + data + "</p>"
	}

	// 渲染 HTML 内容到页面
	return c.HTMLBlob(http.StatusOK, []byte(htmlContent))
}
func main() {
	e := echo.New()

	// 设置模板引擎
	//e.Renderer = initRenderer() // 这里需要实现 initRenderer 函数来设置模板引擎

	// 设置路由
	e.GET("/", getDataHandler)

	// 启动服务器
	e.Start(":8080")

}
