package basic3

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

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
	garde string
}

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	ignored      string         // fields that aren't exported are ignored
}

type Member struct {
	gorm.Model
	Name string
	Age  uint8
}

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	Author
	ID      int
	Upvotes int32
}

type Blog2 struct {
	ID     int64
	Author Author `gorm:"embedded;embeddedPrefix:author_"`
	// Author  Author
	Upvotes int32
}

func Run_1(db *gorm.DB) {
	db.AutoMigrate(&User{})
	// db.AutoMigrate(&Member{})
	// db.AutoMigrate(&Blog{})
	// db.AutoMigrate(&Blog2{})

	user := &User{}
	user.MemberNumber.Valid = true
	db.Create(user)

	// create传指针
	// mem := Member{}
	// db.Create(&mem)
	// fmt.Println(mem.ID)
	// db.Delete(&Member{}, 1)
}
