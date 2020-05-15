package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbName   = "demo"
	user     = "root"
	password = "admin"
	port     = "3306"
	dbDriver = "mysql"
)

type People struct {
	ID     int
	Name   string
	Age    int
	Gender int
}

func ConnectDB() *sql.DB {
	var connectionString = fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s", user, password, port, dbName)
	db, err := sql.Open(dbDriver, connectionString)

	/* Check kết nối */
	if err != nil {
		panic(err)
	}

	return db
}

/* Đọc dữ liệu từ database */
func Read(c *gin.Context) {
	db := ConnectDB()
	defer db.Close()

	query := "SELECT * FROM people"
	result, err := db.Query(query)

	/* Kiểm tra lỗi query */
	if err != nil {
		panic(err)
	}

	/* Nếu không xảy ra lỗi */
	people := []People{}
	for result.Next() {
		p := People{}
		result.Scan(&p.ID, &p.Name, &p.Age, &p.Gender)
		people = append(people, p)
	}

	c.JSON(200, people) /* status OK = 200 */
}

/* Lưu dữ liệu lên database */
func Create(c *gin.Context) {
	/* Đọc dữ liệu được gửi chung request */
	json := People{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	/* Nếu đọc thành công dữ liệu */
	db := ConnectDB()
	defer db.Close()

	query := "INSERT INTO people (name, age, gender) VALUES (?, ?, ?)"
	preparedQuery, _ := db.Prepare(query)
	preparedQuery.Exec(json.Name, json.Age, json.Gender)

	c.JSON(201, gin.H{
		"message": "inserted",
	}) /* status Created = 201 */
}

/* Cập nhật dữ liệu vào database */
func Update(c *gin.Context) {
	/* Đọc dữ liệu được gửi chung request */
	json := People{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	/* Nếu đọc thành công dữ liệu */
	db := ConnectDB()
	defer db.Close()

	query := "UPDATE people SET name = ?, age = ?, gender = ? WHERE id = ?"
	preparedQuery, _ := db.Prepare(query)
	preparedQuery.Exec(json.Name, json.Age, json.Gender, json.ID)

	c.JSON(200, gin.H{
		"message": "updated",
	}) /* status OK = 200 */
}

/* Xóa dòng dữ liệu trên database */
func Delete(c *gin.Context) {
	/* Đọc id từ param trên URL */
	var id = c.Param("id")

	db := ConnectDB()
	defer db.Close()

	query := "DELETE FROM people WHERE id = ?"
	preparedQuery, _ := db.Prepare(query)
	preparedQuery.Exec(id)

	c.JSON(200, gin.H{
		"message": "deleted",
	}) /* status OK = 200 */
}
