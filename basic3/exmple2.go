package basic3

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//事务语句

// 题目1：事务语句
// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表
// （包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，
// 需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，
// 向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

type Account struct {
	ID      uint
	Balance float64
	//定义外键约束 Transaction 表中的转账相关id必须在  Account 中存在
	// 作为转出方的所有交易
	OutgoingTransactions []Transaction `gorm:"foreignKey:From_account_id"`
	// 作为转入方的所有交易
	IncomingTransactions []Transaction `gorm:"foreignKey:To_account_id"`
}

type Transaction struct {
	ID              uint
	From_account_id uint
	To_account_id   uint
	Amount          float64
}

// type Employee struct {
// 	ID         uint
// 	Name       string
// 	Department string
// 	Salary     string
// }

// type Book struct {
// 	ID     int
// 	Title  string
// 	Author string
// 	Price  float64
// }

func Run_2(db *gorm.DB) {
	// db.AutoMigrate(&Book{})
	// db.AutoMigrate(&Transaction{})
	//初始化账户数据
	// accounts := []*Account{
	// 	{Balance: 100},
	// 	{Balance: 200},
	// }
	// db.Debug().Create(accounts)
	// fmt.Println(accounts)
	//生成转账记录，事务处理控制转账的原子性和 转账金额的判断
	err := db.Transaction(func(tx *gorm.DB) error {
		var fromAccount, toAccount Account
		//1.1 检查账户是否存在，并且加锁
		if err := tx.Debug().Model(&Account{}).Clauses(clause.Locking{Strength: "UPDATE"}).First(&fromAccount, 1).Error; err != nil {
			return err
		}
		if err := tx.Debug().Model(&Account{}).Clauses(clause.Locking{Strength: "UPDATE"}).First(&toAccount, 2).Error; err != nil {
			return err
		}
		//1.2转账时判断余额是否充足
		if fromAccount.Balance < 100 {
			return errors.New("转出账户余额不足")
		}
		//1.3 扣减和增加对应账户金额
		fromAccount.Balance -= 100
		toAccount.Balance += 100
		tx.Save(fromAccount)
		tx.Save(toAccount)

		//1.4插入转账记录
		tran := &Transaction{
			From_account_id: 1,
			To_account_id:   2,
			Amount:          100,
		}
		tx.Create(tran)
		return nil
	})
	if err != nil {
		fmt.Println("转账失败：", err)
	}

}
