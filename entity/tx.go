package entity

import (
	"github.com/jinzhu/gorm"
	"tokenup-pay-gateway-vendor-server-demo/model"
	"fmt"
)

type TransactionDb struct {
	Db *gorm.DB `inject:""`
}

func (db *TransactionDb) Save(tx model.Transaction) {
	db.Db.Model(&model.Transaction{}).Create(&tx)
}
func (db *TransactionDb) UpdateStatus(orderId, status string) {
	db.Db.Model(&model.Transaction{}).Where("order_id=?", orderId).Update("status", status)
}
func (db *TransactionDb) FindPage(userId string, pageIndex, pageSize int) []model.Transaction {
	var txs []model.Transaction
	db.Db.Model(&model.Transaction{}).Where("user_id=?", userId).Order("timestamp desc").Offset(pageIndex * pageSize).Limit(pageSize).Find(&txs)
	return txs
}
func (db *TransactionDb) FindBalance(userId string) string {
	var txs []model.Transaction
	db.Db.Model(&model.Transaction{}).Where("user_id=? and status=?", userId, "CONFIRMS").Find(&txs)
	var amount float64
	for _, tx := range txs {
		if tx.Type == 0 {
			amount += tx.Amount
		} else {
			amount -= tx.Amount
		}
	}
	return fmt.Sprint(amount)
}
