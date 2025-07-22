package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	username := "root"
	if os.Getenv("db_username") != "" {
		username = os.Getenv("db_username")
	}
	password := "123456"
	if os.Getenv("db_password") != "" {
		password = os.Getenv("db_password")
	}
	port := "3306"
	if os.Getenv("db_port") != "" {
		port = os.Getenv("db_port")
	}

	// SQL语句练习
	// 题目1:基本CRUD操作
	// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
	// 要求 :
	// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/gorm?charset=utf8mb4&parseTime=True&loc=Local", username, password, port)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("failed to connect database")
		return
	}
	fmt.Println("SQL语句练习-题目1: curd开始执行")
	crudStudents(db)
	// 题目2:事务语句
	// 假设有两个表: accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
	// 要求 :
	// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
	tranMoneyNum := 100.0
	fromAccount := Account{ID: 1, Balance: 500}
	toAccount := Account{ID: 2, Balance: 300}
	fmt.Println("SQL语句练习-题目2: 事务开始执行")
	transferMoney(db, fromAccount, toAccount, tranMoneyNum)

	// Sqlx入门
	// 题目1:使用SQL扩展库进行查询
	// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、
	// department 、 salary 。
	// 要求 :
	// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	// 需要先安装两个依赖：go get -u github.com/jmoiron/sqlx 和 go get -u github.com/go-sql-driver/mysql
	sqlxUrl := fmt.Sprintf("%s:%s@tcp(localhost:%s)/gorm?parseTime=true", username, password, port)
	dbx, err := sqlx.Connect("mysql", sqlxUrl)
	if err != nil {
		panic("failed to connect database with sqlx,ERROR:" + err.Error())
		return
	}
	defer dbx.Close()
	fmt.Println("Sqlx入门-题目1: 使用SQL扩展库进行查询开始执行")
	sqlxExtendControl(dbx)
	// 题目2:实现类型安全映射
	// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
	// 要求 :
	// 定义一个 Book 结构体，包含与 books 表对应的字段。
	// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，
	// 并将结果映射到 Book 结构体切片中，确保类型安全。
	fmt.Println("Sqlx入门-题目2: 实现类型安全映射开始执行")
	price := 50.0
	getPriceGreaterThanXPrice(dbx, price)

	// 进阶gorm
	// 题目1：模型定义
	// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
	// 要求 ：
	// 使用Gorm定义 User 、 Post 和 Comment 模型，
	// 其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
	// Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
	// 编写Go代码，使用Gorm创建这些模型对应的数据库表。
	fmt.Println("进阶gorm-题目1: 模型定义开始执行")
	userCount := 10
	// 受前面sqlx的影响，db需要重新创建
	db2, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("failed to connect database")
		return
	}
	generateDefaultData(db2, userCount)
	// 题目2：关联查询
	// 基于上述博客系统的模型定义。
	// 要求 ：
	// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
	fmt.Println("进阶gorm-题目2: 关联查询开始执行")
	userId := 1 // 假设查询用户ID为1的用户
	associationSearch(db2, userId)
	// 题目3：钩子函数
	// 继续使用博客系统的模型。
	// 要求 ：
	// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	fmt.Println("进阶gorm-题目3: 钩子函数开始执行")
	db2.Create(&Post{UserId: 1, Title: "test1"})
	db2.Create(&Comment{PostId: 1, Content: "test"})
}

type User struct {
	Id       int    `gorm:"primaryKey"`
	Username string `gorm:"not null"`
	PostNum  int    `gorm:"comment:文章数量统计;default:0"`
	Posts    []Post `gorm:"foreignKey:UserId;references:Id"`
}
type Post struct {
	Id         int       `gorm:"primaryKey"`
	UserId     int       `gorm:"not null"`
	Title      string    `gorm:"not null"`
	HasComment int       `grom:"comment: 文章是否被评论字段0-未评论,1-评论;default:0"`
	Comments   []Comment `gorm:"foreignKey:PostId;references:Id"`
	// User     User      `gorm:"foreignKey:UserId"` // 添加反向引用
}
type Comment struct {
	Id      int    //`gorm:"primaryKey"`
	PostId  int    //`gorm:"not null;index"`
	Content string //`gorm:"not null"`
}

func (post *Post) AfterCreate(tx *gorm.DB) (err error) {
	posts := []Post{}
	tx.Where("user_id = ?", post.UserId).Find(&posts)
	tx.Model(&User{}).Where("id = ?", post.UserId).Update("post_num", len(posts))
	return
}
func (comment *Comment) AfterCreate(tx *gorm.DB) (err error) {
	comments := []Comment{}
	tx.Where("post_id = ?", comment.PostId).Find(&comments)
	hasComment := 0
	if len(comments) > 0 {
		hasComment = 1
	}
	tx.Model(&Post{}).Where("id = ?", comment.PostId).Update("has_comment", hasComment)
	return
}

func associationSearch(db *gorm.DB, userId int) {
	users := []User{}
	db.Where("id = ?", userId).Preload("Posts.Comments").Find(&users)
	fmt.Println("查询用户ID为", userId, "的用户发布的所有文章及其对应的评论信息:", users)

	post := Post{}
	// db.Raw("SELECT posts.*, Count( comments.id ) AS comment_count from posts	LEFT JOIN comments ON comments.post_id = posts.id  GROUP BY	posts.id  ORDER BY	comment_count DESC 	LIMIT 1").Scan(&post)

	db.Select("posts.*, COUNT(comments.id) AS comment_count").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count DESC").First(&post)

	var commentsNum int64 = int64(0)
	db.Model(&Comment{}).Where("post_id = ?", post.Id).Count(&commentsNum)
	fmt.Println("查询评论数量最多的文章信息:", post, "评论数：", commentsNum)
}

func generateDefaultData(db *gorm.DB, userCount int) {
	// 自动迁移模型：有先后依赖关系，按依赖关系顺序迁移
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&User{})
	users := []User{}
	for i := 1; i <= userCount; i++ {
		user := User{
			Id:       i,
			Username: fmt.Sprintf("用户%d", i),
			Posts: []Post{
				{
					Id:     i,
					UserId: i,
					Title:  fmt.Sprintf("文章%d", i),
					Comments: []Comment{
						{
							Id:      i,
							PostId:  i,
							Content: fmt.Sprintf("评论%d", i),
						},
						{
							Id:      i + 1,
							PostId:  i,
							Content: fmt.Sprintf("评论%d", i+1),
						},
						{
							Id:      i + 2,
							PostId:  i,
							Content: fmt.Sprintf("评论%d", i+2),
						},
					},
				},
				{
					Id:     i + 10,
					UserId: i,
					Title:  fmt.Sprintf("文章%d", i),
					Comments: []Comment{
						{
							Id:      i + 10,
							PostId:  i + 10,
							Content: fmt.Sprintf("评论%d", i+10),
						},
						{
							Id:      i + 11,
							PostId:  i + 10,
							Content: fmt.Sprintf("评论%d", i+11),
						},
						{
							Id:      i + 12,
							PostId:  i + 10,
							Content: fmt.Sprintf("评论%d", i+12),
						},
					},
				},
			},
		}
		users = append(users, user)
	}
	if err := db.Create(&users).Error; err != nil {
		fmt.Printf("插入用户%d数据失败: %v\n", userCount, err)
		return
	}
	fmt.Printf("插入用户%d数据成功\n", userCount)

}

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func getPriceGreaterThanXPrice(dbx *sqlx.DB, price float64) {
	// 创建表
	dbx.MustExec(`	
		CREATE TABLE IF NOT EXISTS book (
			id INT PRIMARY KEY,
			title VARCHAR(100) NOT NULL,
			author VARCHAR(50) NOT NULL,
			price DECIMAL(10, 2) NOT NULL
		)
	`)
	// 插入数据
	count := 10
	batchBooks := []Book{}
	for i := 1; i <= count; i++ {
		book := Book{
			ID:     i,
			Title:  fmt.Sprintf("书籍%d", i),
			Author: fmt.Sprintf("作者%d", i),
			Price:  float64(10 + i*5), // 模拟价格从10到60
		}
		batchBooks = append(batchBooks, book)
	}
	_, err := dbx.NamedExec(`
		INSERT INTO book (id, title, author, price)
		VALUES (:id, :title, :author, :price)
		ON DUPLICATE KEY UPDATE
			title = VALUES(title),
			author = VALUES(author),
			price = VALUES(price)
	`, batchBooks)
	if err != nil {
		fmt.Println("插入书籍信息失败")
		return
	}
	// 查询价格大于指定值的书籍
	books := []Book{}
	if dbx.Select(&books, "SELECT * FROM book WHERE price > ?", price) != nil {
		fmt.Println("查询价格大于指定值的书籍信息失败")
		return
	}
	fmt.Println("查询价格大于", price, "的书籍信息:", books)
}

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func sqlxExtendControl(dbx *sqlx.DB) {
	// dbx.AutoMigrate(&Employee{})
	// 创建表
	dbx.MustExec(`
		CREATE TABLE IF NOT EXISTS employees (
			id INT PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			department VARCHAR(50) NOT NULL,
			salary INT NOT NULL
		)
	`)

	// 模拟插入10条员工数据
	count := 10
	batchEmployees := []Employee{}
	for i := 1; i <= count; i++ {
		employee := Employee{
			ID:         i,
			Name:       fmt.Sprintf("员工%d", i),
			Department: "技术部",
			Salary:     float64(3000 + i*100), // 模拟工资从3100到4000
		}
		batchEmployees = append(batchEmployees, employee)
	}
	_, err := dbx.NamedExec(`
		INSERT INTO employees (id, name, department, salary)
		VALUES (:id, :name, :department, :salary)
		ON DUPLICATE KEY UPDATE
			name = VALUES(name),
			department = VALUES(department),
			salary = VALUES(salary)
	`, batchEmployees)
	if err != nil {
		fmt.Println("插入员工信息失败")
		return
	}

	// 1.查询所有部门为"技术部"的员工信息
	employees := []Employee{}
	if dbx.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部") != nil {
		fmt.Println("查询部门为技术部的员工信息失败")
		return
	}
	fmt.Println("查询部门为技术部的员工信息:", employees)

	// 2.查询工资最高的员工信息
	topSalaryEmployee := Employee{}
	if dbx.Get(&topSalaryEmployee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1") != nil {
		fmt.Println("查询工资最高的员工信息失败")
		return
	}
	fmt.Println("查询工资最高的员工信息:", topSalaryEmployee)

}

type Transfer interface {
	isEnougn(tranMoneyNum float64) bool
}
type Account struct {
	ID      int
	Balance float64
}
type Transaction struct {
	ID            int
	FromAccountID int
	ToAccountID   int
	Amount        float64
}

func (a *Account) isEnougn(tranMoneyNum float64) bool {
	if a.Balance >= tranMoneyNum {
		return true
	}
	return false
}
func transferMoney(db *gorm.DB, fromAccount Account, toAccount Account, amount float64) {
	db.AutoMigrate(&fromAccount, &toAccount, &Transaction{})
	db.Begin()
	if !fromAccount.isEnougn(amount) {
		// 资金不足直接回滚
		db.Rollback()
		fmt.Println("转账失败:账户余额不足")
		return
	}
	fromAccount.Balance -= amount
	toAccount.Balance += amount
	if db.Save(&fromAccount).Error != nil {
		db.Rollback()
		fmt.Println("转账失败:更新转出账户失败")
		return
	}
	if db.Save(&toAccount).Error != nil {
		db.Rollback()
		fmt.Println("转账失败:更新转入账户失败")
		return
	}
	transaction := Transaction{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        amount,
	}
	if db.Create(&transaction).Error != nil {
		db.Rollback()
		fmt.Println("转账失败:记录转账信息失败")
		return
	}
	db.Commit()
	fmt.Println("转账成功:从账户", fromAccount.ID, "向账户", toAccount.ID, "转账金额", amount)
}

type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

func crudStudents(db *gorm.DB) {
	db.AutoMigrate(&Student{})
	// 1.插入学生信息:Name: "张三", Age: 20, Grade: "三年级"
	student := Student{Name: "张三", Age: 20, Grade: "三年级"}
	if db.Create(&student).Error != nil {
		fmt.Println("插入学生信息失败")
	} else {
		fmt.Println("插入学生信息成功:", student)
	}
	// 2.查询年龄大于18岁的学生信息
	students := []Student{}
	if db.Where("age > ?", 18).Find(&students).Error != nil {
		fmt.Println("查询年龄大于18岁的学生信息失败")
	} else {
		fmt.Println("查询年龄大于18岁的学生信息:", students)
	}
	// 3.更新姓名为"张三"的学生年级为"四年级"。Model是用于指定表，如果没有Find\Delete等指定操作结构体即表，则需要指定
	if db.Debug().Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级").Error != nil {
		fmt.Println("更新学生年级失败")
	} else {
		fmt.Println("更新学生年级成功")
	}
	// 4.删除年龄小于15的学生
	if db.Where("age < ?", 15).Delete(&Student{}).Error != nil {
		fmt.Println("删除年龄小于15岁的学生记录失败")
	} else {
		fmt.Println("删除年龄小于15岁的学生记录成功")
	}
}
