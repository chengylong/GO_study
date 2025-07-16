package basic3

import (
	"gorm.io/gorm"
)

//进阶gorm

// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。

type User struct {
	ID         uint
	Name       string
	Post_count uint
	//定义外键约束 Post 表中的文章id必须在  User 中存在
	//关联用户的文章
	OutgoingPosts []Post `gorm:"foreignKey:User_id"`
}

type Post struct {
	ID             uint
	Title          string
	User_id        uint
	Comment_count  uint
	Comment_status string
	//关联文章的评论
	OutgoingComments []Comment `gorm:"foreignKey:Post_id"`
}

type Comment struct {
	ID      uint
	Content string
	Post_id uint
}

type CountRes struct {
	Post_id uint
	Count   uint
}

func Run_4(db *gorm.DB) {
	// 题目1：模型定义
	// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
	// 要求 ：
	// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
	// 编写Go代码，使用Gorm创建这些模型对应的数据库表。
	//初始化表结构
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&Post{})
	// db.AutoMigrate(&Comment{})

	// 	题目2：关联查询
	// 基于上述博客系统的模型定义。
	// 要求 ：
	// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
	//初始化数据
	// users := []User{{Name: "张三"}, {Name: "李四"}, {Name: "王五"}}
	// db.Create(users)

	// posts := []Post{
	// 	//测试外键约束
	// 	// {Title: "标题1", User_id: 5},
	// 	//插入数据
	// 	{Title: "标题2", User_id: 1},
	// 	{Title: "标题3", User_id: 1},
	// 	{Title: "标题4", User_id: 2},
	// 	{Title: "标题5", User_id: 2},
	// 	{Title: "标题6", User_id: 3},
	// }
	// db.Create(posts)

	// comments := []Comment{
	// 	//测试外键约束
	// 	// {Content: "评论1", Post_id: 100},
	// 	//插入数据
	// 	{Content: "评论2", Post_id: 2},
	// 	{Content: "评论3", Post_id: 2},
	// 	{Content: "评论4", Post_id: 3},
	// 	{Content: "评论5", Post_id: 3},
	// 	{Content: "评论6", Post_id: 2},
	// }
	// db.Create(comments)

	//题目1 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	//使用预加载，查出用户 张三关联的文章以及文章的评论
	// var user User
	// db.Debug().Preload("OutgoingPosts.OutgoingComments").First(&user, 1)
	// fmt.Println(user)
	// // 输出所有评论内容
	// for _, post := range user.OutgoingPosts {
	// 	for _, comment := range post.OutgoingComments {
	// 		fmt.Println("文章：", post.Title, "评论：", comment.Content)
	// 	}
	// }

	//题目2 编写Go代码，使用Gorm查询评论数量最多的文章信息。
	//2.1 聚合查询查出评论最多的文章ID
	// res := []CountRes{}
	// db.Debug().Model(&Comment{}).Select("post_id,count(1) as count").Group("post_id").Order("count desc").Limit(1).Scan(&res)
	// fmt.Println(res)

	// //2.2 根据得到的结果，查出文章信息
	// var post Post
	// //预加载
	// db.Debug().Preload("OutgoingComments").First(&post, res[0].Post_id)
	// fmt.Println(post)

	// //查询文章
	// db.Debug().First(&post, res[0].Post_id)
	// fmt.Println(post)

	// 	题目3：钩子函数
	// 继续使用博客系统的模型。
	// 要求 ：
	//3.1 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	//创建文章

	// posts := []Post{
	// 	{Title: "哈哈", User_id: 1},
	// 	{Title: "嘻嘻", User_id: 1},
	// }
	// db.Create(&posts)
	// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	//先查出来再删除,否则comment 的 post_id无赋值，钩子函数需要用到该值，会出问题
	var comment Comment
	db.Debug().First(&comment, 10)
	db.Delete(&comment)

}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	//文章创建后 自动更新用户的 PostCount 字段
	err = tx.Debug().Model(&User{}).Where("id= ?", p.User_id).Update("post_count", gorm.Expr("post_count +?", 1)).Error
	return err
}

func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	//评论创建后 自动更新文章的 评论数量 comment_count 字段
	err = tx.Debug().Model(&Post{}).Where("id= ?", c.Post_id).Update("comment_count", gorm.Expr("comment_count +?", 1)).Error
	return err
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	//在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	// 1. 先将评论数减一
	err = tx.Debug().Model(&Post{}).Where("id = ?", c.Post_id).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error
	//避免 评论数被减为负数
	// err = tx.Model(&Post{}).Where("id = ?", c.Post_id).Update("comment_count", gorm.Expr("GREATEST(comment_count - 1, 0)")).Error
	if err != nil {
		return err
	}

	//查询该文章最新评论数
	var post Post
	err = tx.Debug().Model(&Post{}).Select("comment_count").Where("id = ?", c.Post_id).First(&post).Error
	if err != nil {
		return err
	}
	//判断评论数
	if post.Comment_count == 0 {
		err = tx.Debug().Model(&Post{}).Where("id = ?", c.Post_id).Update("Comment_status", "无评论").Error
	}
	return err
}
