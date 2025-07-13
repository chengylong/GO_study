package main

import (
	"fmt"
	"time"
)

//1，指针

// 1.1：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
// 在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func editInt(num *int) {
	*num = *num + 1
	//新定义的指针变量和原指针没有关系
	// fmt.Println("p:", p)
	// num = &p
}

// 1.2:实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 1.2.1切片指针
func multiply2_1(nums *[]int) {
	for i, _ := range *nums {
		(*nums)[i] *= 2
	}
}

// 1.2.2 包含指针的切片
func multiply2_2(nums []*int) {
	for _, v := range nums {
		*v *= 2
	}
}

//2. Goroutine

// 2.1 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func goroutine_test() {
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Println(i)
			}
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	}()
}

// 2.2设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 任务执行结果结构
type TaskResult struct {
	ID       int
	Duration time.Duration
}

// 任务类型定义
type Task func()

//简单time实现
// func Scheduler(tasks []Task) []TaskResult {
// 	// var wg sync.WaitGroup
// 	results := make([]TaskResult, len(tasks))
// 	for i, task := range tasks {
// 		go func(id int, t Task) {
// 			start := time.Now()
// 			t()
// 			duration := time.Since(start) // 计算耗时
// 			results[i] = TaskResult{
// 				ID:       i,
// 				Duration: duration,
// 			}

// 		}(i, task)
// 	}
// 	time.Sleep(3 * time.Second)
// 	return results
// }

// 使用 WaitGroup 实现
// func Scheduler(tasks []Task) []TaskResult {
// 	var wg sync.WaitGroup
// 	results := make([]TaskResult, len(tasks))
// 	for i, task := range tasks {
// 		//加入1个协程处理
// 		wg.Add(1)
// 		go func(id int, t Task) {
// 			//协程处理结束后，释放出协程 类似于 wg.Add(-1)
// 			defer wg.Done()
// 			start := time.Now()
// 			t()
// 			duration := time.Since(start) // 计算耗时
// 			results[i] = TaskResult{
// 				ID:       i,
// 				Duration: duration,
// 			}

// 		}(i, task)
// 		//等待 wg中所有加入的协程处理结束
// 		wg.Wait()
// 	}
// 	// time.Sleep(3 * time.Second)
// 	return results
// }
// func main() {
// int1 := 10
// fmt.Println(int1)
// editInt(&int1)
// fmt.Println(int1)
// var nums = []int{4, 1, 2, 1, 2}
// nums2 := make([]*int, 5)
// for i := range nums2 {
// 	nums2[i] = new(int)
// 	*nums2[i] = []int{4, 1, 2, 1, 2}[i]
// }
// multiply2(nums2)
// for _, v := range nums2 {
// 	fmt.Println(*v)
// }
// fmt.Println(nums2)
// goroutine_test()
// time.Sleep(1 * time.Second)
// 示例任务定义
// tasks := []Task{
// 	func() { time.Sleep(300 * time.Millisecond) },
// 	func() { time.Sleep(500 * time.Millisecond) },
// 	func() { time.Sleep(200 * time.Millisecond) },
// }
// // 执行调度
// results := Scheduler(tasks)

// // 打印结果
// fmt.Println("任务执行统计:")
// for _, r := range results {
// 	fmt.Printf("任务 %d 耗时: %v", r.ID, r.Duration)
// }

// }

//3. 面向对象
// 3.1定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，
// 创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。

// type Shape interface {
// 	Area()
// 	Perimeter()
// }
// type Rectangle struct {
// 	long float32
// 	wide float32
// }
// type Circle struct {
// 	Radius float32
// }

// // 实现接口重写方法
// // 长方形
// func (r Rectangle) Area() float32 {
// 	return r.long * r.wide
// }

// func (r Rectangle) Perimeter() float32 {
// 	return 2 * (r.long + r.wide)
// }

// // 圆
// func (r Circle) Area() float32 {
// 	return math.Pi * (r.Radius * r.Radius)
// }

// func (r Circle) Perimeter() float32 {
// 	return 2 * math.Pi * r.Radius
// }

// func main() {
// 	r := Rectangle{
// 		long: 2.0,
// 		wide: 3.0,
// 	}
// 	fmt.Println("长方形面积：", r.Area())
// 	fmt.Println("长方形周长：", r.Perimeter())
// 	c := Circle{
// 		Radius: 3,
// 	}
// 	fmt.Println("圆形面积：", c.Area())
// 	fmt.Println("圆形周长：", c.Perimeter())

// }

// 3.2 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
// 组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Println("名字：", e.Name)
	fmt.Println("年龄:", e.Age)
	fmt.Println("编号:", e.EmployeeID)
}

func main() {
	person1 := Employee{
		Person: Person{
			Name: "张三",
			Age:  20,
		},

		EmployeeID: "001",
	}
	person2 := Employee{
		Person: Person{
			Name: "李四",
			Age:  20,
		},

		EmployeeID: "001",
	}
	person1.PrintInfo()
	person2.PrintInfo()

}
