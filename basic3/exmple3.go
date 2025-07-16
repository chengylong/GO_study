package basic3

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

//Sqlx入门

// 题目1：使用SQL扩展库进行查询
// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

// type Employee struct {
// 	ID         uint
// 	Name       string
// 	Department string
// 	Salary     float64
// }

type Book struct {
	ID     int
	Title  string
	Author string
	Price  float64
}

// 题目2：实现类型安全映射
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

func Run_3(db *sqlx.DB) {
	//初始化数据
	// result, err := db.Exec("insert into Employees (name,department,salary) values (?,?,?)", "张三", "技术部", 10000)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// id, _ := result.LastInsertId()
	// fmt.Println("新用户ID:", id)

	// result, err := db.Exec("insert into Employees (name,department,salary) values (?,?,?)", "王五", "技术部", "8000")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// id, _ := result.LastInsertId()
	// fmt.Println("新用户ID:", id)

	// res2, err2 := db.Exec("insert into Employees (name,department,salary) values (?,?,?)", "李四", "市场部", "15000")
	// if err2 != nil {
	// 	log.Fatalln(err)
	// }
	// id2, _ := res2.LastInsertId()
	// fmt.Println("新用户ID:", id2)

	// 2.1 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	// var employees []Employee
	// err := db.Select(&employees, "select * from employees where department =?", "技术部")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(employees)

	// 2.2 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	// var employee Employee
	// err := db.Get(&employee, "select * from employees where 1=1 order by salary desc")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(employee)

	// type Book struct {
	// 	ID     int
	// 	Title  string
	// 	Author string
	// 	Price  float64
	// }
	// 	题目2：实现类型安全映射
	// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
	// 要求 ：
	// 定义一个 Book 结构体，包含与 books 表对应的字段。
	// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

	//初始化数据
	// result, err := db.Exec("insert into books (Title,Author,Price) values (?,?,?)", "福尔摩斯探案集", "au3", 10)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// id, _ := result.LastInsertId()
	// fmt.Println("书:", id)
	var books []Book
	err := db.Select(&books, "select * from books where price > ?", 50)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(books)

}
