package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/banking/dto"
	"github.com/banking/service"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah *AccountHandler) NewAccount(c *gin.Context) {

	customerId := c.Params.ByName("id")
	var request dto.NewAccountRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	request.CustomerId = customerId
	fmt.Println(request)
	account, appErr := ah.service.NewAccount(request)
	if appErr != nil {
		c.JSON(appErr.Code, gin.H{
			"err": appErr.Message,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": account,
		})
	}
}

func (ah *AccountHandler) MakeTransaction(c *gin.Context) {
	fmt.Println("inside Handler")
	customerId := c.Params.ByName("id")
	accountId := c.Params.ByName("accountId")

	var request dto.TransactionRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	} else {
		request.AccountId = accountId
		request.CustomerId = customerId
		log.Println("Reached Here")
		account, appError := ah.service.MakeTransaction(request)
		log.Printf("Resp :- %v", account)
		if appError != nil {
			c.JSON(appError.Code, gin.H{
				"error": appError.Message,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data": account,
			})
		}
	}

}
