package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()
	router.GET("/balance/", getBalance)
	router.GET("/deposit/:input", deposit)
	router.GET("/withdraw/:input", withdraw)
	router.Run(":8080")
}

// 提款
func withdraw(c *gin.Context) {
	input := c.Param("input")
	amount, err := strconv.Atoi(input)
	if err != nil {
		preview(c, bank, err)
	}
	err = bank.Withdrawal(amount)
	if err != nil {
		preview(c, bank, err)
	}
}

// 存款
func deposit(c *gin.Context) {
	input := c.Param("input")
	amount, err := strconv.Atoi(input)
	if err != nil {
		preview(c, bank, err)
	}
	err = bank.Save(amount)
	if err != nil {
		preview(c, bank, err)
	}

}

// 取得餘額
func getBalance(c *gin.Context) {
	preview(c, bank, nil)
}

func preview(c *gin.Context, b *Bank, err error) {
	r := struct {
		Amount  int    `json="amount"`
		Status  string `json="status"`
		Message string `json="message"`
	}{
		Amount:  b.GetAmount(),
		Status:  "ok",
		Message: "",
	}
	if err != nil {
		r.Amount = 0
		r.Status = "fail"
		r.Message = err.Error()
	}

	c.JSON(http.StatusOK, r)
}

type Bank struct {
	amount int `json="amount"`
}

var bank = &Bank{
	amount: 0,
}

func (b *Bank) GetAmount() int {
	return b.amount
}

func (b *Bank) Save(amount int) (err error) {
	if amount < 0 {
		return errors.New("操作失敗，存款金額須大於０")
	}
	b.amount += amount
	return
}

func (b *Bank) Withdrawal(amount int) (err error) {
	if amount < 0 {
		return errors.New("操作失敗，提款金額須大於０")
	}
	if b.amount-amount < 0 {
		return errors.New("操作失敗，存款不足")
	}
	b.amount -= amount
	return
}
