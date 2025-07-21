package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// 指针
	// 题目1 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
	// 考察点 ：指针的使用、值传递与引用传递的区别。
	num := 2
	fmt.Print("题目1:原先的值为:", num)
	incrementTenWithPointer(&num)
	fmt.Printf(" 修改后的值为:%d\n", num)
	// 题目2 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
	// 考察点 ：指针运算、切片操作。
	numSlices := []int{1, 2, 3, 4, 5}
	fmt.Print("题目2:原先的切片为:", numSlices)
	multiplySliceByTwo(&numSlices)
	fmt.Printf(" 修改后的切片为:%v\n", numSlices)

	// Goroutine
	// 题目3 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
	// 考察点 ： go 关键字的使用、协程的并发执行。
	goRoutineTest()
	// time.Sleep(1 * time.Second) // 等待协程执行完毕
	// 题目4 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
	// 考察点 ：协程原理、并发任务调度。
	taskScheduler := &TaskScheduler{}
	taskScheduler.AddTask(1, func() {
		time.Sleep(1 * time.Second)
		fmt.Println("任务1完成")
	})
	taskScheduler.AddTask(2, func() { time.Sleep(2 * time.Second) })
	taskScheduler.AddTask(3, func() { time.Sleep(3 * time.Second) })
	taskScheduler.Run()

	// 面向对象
	// 题目5 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
	// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
	// 考察点 ：接口的定义与实现、面向对象编程风格。
	rectangle := &Rectangle{width: 3, height: 4}
	circle := &Circle{radius: 5}
	fmt.Println("题目5:矩形和圆的面积与周长:")
	fmt.Printf(">>>矩形的面积: %.2f, 周长: %.2f\n", rectangle.Area(), rectangle.Perimeter())
	fmt.Printf(">>>圆的面积: %.2f, 周长: %.2f\n", circle.Area(), circle.Perimeter())
	// 题目6 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
	// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
	// 考察点 ：组合的使用、方法接收者。
	employee := &Employee{
		Person: Person{
			Name: "张三",
			Age:  30,
		},
		EmployeeID: "E12345",
	}
	fmt.Println("题目6:员工信息:")
	employee.PrintInfo()

	// Channel
	// 题目7 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，
	// 另一个协程从通道中接收这些整数并打印出来。
	// 考察点 ：通道的基本使用、协程间通信。
	fmt.Print("题目7:协程间通信：接收到的信息如下:")
	channelTransferTest()
	// 题目8 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
	// 考察点 ：通道的缓冲机制。
	fmt.Print("题目8:协程间(缓冲)通信:接收到的信息如下：")
	channelTransferBufferTest()

	// 锁机制
	// 题目9 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	// 考察点 ： sync.Mutex 的使用、并发数据安全。
	shareValue := 10
	corontineCount := 10
	incrementCount := 1000
	plusByGoRoutineBySafe(&shareValue, corontineCount, incrementCount)
	fmt.Printf("题目9:使用sync.Mutex保护的计数器的值为:%d\n", shareValue)
	// 题目10 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	// 考察点 ：原子操作、并发数据安全。
	var shareOrigin int64 = 10
	shareAutomic := atomic.LoadInt64(&shareOrigin)
	plusByAutomic(&shareAutomic, corontineCount, incrementCount)
	fmt.Printf("题目10:使用原子操作的计数器的值为:%d\n", shareAutomic)
}

func plusByAutomic(shareValue *int64, corontineCount int, incrementCount int) {
	var wg sync.WaitGroup
	wg.Add(corontineCount)
	for i := 0; i < corontineCount; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementCount; j++ {
				atomic.AddInt64(shareValue, 1) // 使用原子操作递增共享资源
			}
		}()
	}
	wg.Wait()
}

func plusByGoRoutineBySafe(shareValue *int, corontineCount int, incrementCount int) {
	var wg sync.WaitGroup
	var mu sync.Mutex      // 创建一个互斥锁
	wg.Add(corontineCount) // 设置等待组计数为协程数量
	for i := 0; i < corontineCount; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementCount; j++ {
				mu.Lock()        // 上锁，保护共享资源
				*shareValue += 1 // 对共享资源进行递增操作
				mu.Unlock()      // 解锁
			}
		}()
	}
	wg.Wait()
}

func channelTransferBufferTest() {
	ch := make(chan int, 50)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		// 生成从1到100的整数并发送到通道
		for i := 1; i <= 100; i++ {
			ch <- i // 发送整数到通道
		}
		close(ch) // 关闭通道
	}()
	go func() {
		defer wg.Done()
		// 从通道接收整数
		for num := range ch {
			fmt.Printf("%d,", num)
		}
	}()
	wg.Wait() // 等待所有协程完成
	fmt.Println()
}

func channelTransferTest() {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		// 生成从1到10的整数并发送到通道
		for i := 1; i <= 10; i++ {
			ch <- i // 发送整数到通道
		}
		close(ch) // 关闭通道
	}()
	go func() {
		defer wg.Done()
		// 从通道接收整数
		for num := range ch {
			fmt.Printf("%d,", num)
		}
	}()
	wg.Wait() // 等待所有协程完成
	fmt.Println()
}

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	EmployeeID string
}

func (e *Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %s\n", e.Name, e.Age, e.EmployeeID)
}

type Shape interface {
	Area() float64      // 面积
	Perimeter() float64 // 周长
}
type Rectangle struct {
	width  float64
	height float64
}
type Circle struct {
	radius float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}
func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius // 使用近似值π=3.14
}
func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius // 使用近似值π=3.14
}

type Task struct {
	TaskId  int
	Action  func()
	StartT  time.Time
	EndT    time.Time
	Elapsed time.Duration
}
type TaskScheduler struct {
	tasks []Task
	wg    sync.WaitGroup
}

func (ts *TaskScheduler) AddTask(taskId int, action func()) {
	task := Task{
		TaskId: taskId,
		Action: action,
	}
	ts.tasks = append(ts.tasks, task)
}
func (ts *TaskScheduler) Run() {
	ts.wg.Add(len(ts.tasks)) // 设置等待组计数为任务数量
	for _, task := range ts.tasks {
		go func(task Task) {
			defer ts.wg.Done()
			task.StartT = time.Now()               // 记录开始时间
			task.Action()                          // 执行任务
			task.EndT = time.Now()                 // 记录结束时间
			task.Elapsed = time.Since(task.StartT) // 计算执行时间
			// fmt.Printf("任务Id:%d, 执行时间:%v,开始时间:%v,结束时间:%v\n", task.TaskId, task.EndT.Second()-task.StartT.Second(), task.StartT, task.EndT)
		}(task)
	}
	ts.wg.Wait() // 等待所有任务完成
	fmt.Println("题目四：结果如下：")
	for _, task := range ts.tasks {
		fmt.Printf("任务Id:%d, 执行时间:%v\n", task.TaskId, task.Elapsed.Seconds())
	}
}

func goRoutineTest() {
	var wg sync.WaitGroup
	wg.Add(2) // 添加两个协程的计数

	fmt.Print("题目3:")
	go func() {
		defer wg.Done() // 完成时通知waitGroup
		for i := 1; i <= 10; i += 2 {
			fmt.Print("奇数:", i, "  ")
		}
	}()
	go func() {
		defer wg.Done() // 完成时通知waitGroup
		for i := 2; i <= 10; i += 2 {
			fmt.Print("偶数:", i, "  ")
		}
	}()

	wg.Wait() // 等待所有协程完成
}

func incrementTenWithPointer(num *int) {
	*num += 10
}

func multiplySliceByTwo(numSlices *[]int) {
	for i := range *numSlices {
		(*numSlices)[i] *= 2
	}
}
