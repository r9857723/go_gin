package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var balance = 0

type Result struct {
	Amount int `json="amount"`
	Status string `json="status"`
	Message string `json="message"`
}

func main() {
	router := gin.Default()
	router.GET("/balance/", getBalance)
	router.GET("/deposit/:input", deposit)
	router.GET("/withdraw/:input", withdraw)
	router.Run(":8080")
}

func wrapResponse(context *gin.Context, amount int ,err error) {
	var result = Result{
		Amount: balance,
		Status: "ok",
		Message: "",
	}
	if err != nil {
		result.Amount = 0
		result.Status = "failed"
		result.Message = err.Error()
	}
	context.JSON(http.StatusOK, result)
}

func withdraw(c *gin.Context) {
	var status string
	var msg string
	input := c.Param("input")
	amount, err := strconv.Atoi(input)
	if err == nil {
		if amount <= 0 {
			amount = 0
			status = "failed"
			msg = "提款金額小於等於０ 提款失敗"
		} else {
			if balance - amount < 0 {
				amount = 0
				status = "failed"
				msg = "餘額不足 提款失敗"
			} else {
				balance -= amount
				status = "ok"
				msg = "提款成功"
			}
		}
	} else {
		amount = 0
		status = "failed"
		msg = "提款失敗"
	}

	c.JSON(http.StatusOK, gin.H{
		"amount":  amount,
		"status":  status,
		"message": msg,
	})
}

func deposit(c *gin.Context) {
	var status string
	var msg string
	input := c.Param("input")
	amount, err := strconv.Atoi(input)
	if err == nil {
		if amount <= 0 {
			amount = 0
			status = "failed"
			msg = "存款金額小於0,存款失敗"
		} else {
			balance += amount
			status = "ok"
			msg = "成功存款" + strconv.Itoa(amount) + "元"
		}
	} else {
		amount = 0
		status = "failed"
		msg = "存款失敗"
	}

	c.JSON(http.StatusOK, gin.H{
		"amount":  amount,
		"status":  status,
		"message": msg,
	})
}
func getBalance(c *gin.Context) {
	msg := "您的帳戶內有:" + strconv.Itoa(balance) + "元"
	c.JSON(http.StatusOK, gin.H{
		"amount":  balance,
		"status":  "OK",
		"message": msg,
	})
}
