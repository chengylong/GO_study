package basic3

import (
	"gorm.io/gorm"
)

//gorm crud 操作

// 题目1：基本CRUD操作
// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
// 要求 ：
// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

type Student struct {
	ID    uint
	Name  string
	Age   uint
	Grade string
}

func Run_1(db *gorm.DB) {
	// db.AutoMigrate(&Student{})
	//1.1编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	// stu := &Student{
	// 	Name:  "王五",
	// 	Age:   20,
	// 	Grade: "三年级",
	// }
	// stu := &Student{
	// 	Name:  "李四",
	// 	Age:   17,
	// 	Grade: "三年级",
	// }
	// db.Create(stu)
	// stus := []*Student{
	// 	{
	// 		Name:  "张三",
	// 		Age:   20,
	// 		Grade: "三年级"},

	// 	{
	// 		Name:  "李四",
	// 		Age:   17,
	// 		Grade: "三年级"},
	// 	{
	// 		Name:  "王五",
	// 		Age:   21,
	// 		Grade: "三年级"},
	// }

	// db.Create(stus)
	//1.2编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	// stus := []Student{}
	// db.Where("age>?", 18).First(&stus)
	// db.Where("age>?", 18).Find(&stus)
	// fmt.Println(stus)
	//1.3编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	// db.Model(&Student{}).Where("name=?", "张三").Update("grade", "四年级")
	//1.4编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	// db.Debug().Model(&Student{}).Where("age<?", 15).Delete(&Student{})

}
