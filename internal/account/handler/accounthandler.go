package handler

import (
	"FinanceApp/internal/account"
	"FinanceApp/internal/account/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AccountHandler struct {
	accountService services.AccountService
}

func NewAccountHandler(accService services.AccountService, router *gin.RouterGroup) *AccountHandler {
	return &AccountHandler{accountService: accService}
}

func (a *AccountHandler) RegisterAccountRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.Group("/accounts")
	routerGroup.POST("", a.Create)
	routerGroup.PUT("/:id", a.Update)
	routerGroup.DELETE("/:id", a.Delete)
	routerGroup.GET("/:id", a.Get)
}

func (a *AccountHandler) Create(ctx *gin.Context) {
	var accreq account.CreateAccountRequest
	reqContext := ctx.Request.Context()
	if err := ctx.ShouldBindJSON(&accreq); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": validationErrors.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := a.accountService.Create(reqContext, accreq)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, account)

}

func (a *AccountHandler) Get(ctx *gin.Context) {

	//TODO

	//	 a.accountService.Get(c.Request.Context(), account.AccountUriRequest)

}

func (a *AccountHandler) Update(ctx *gin.Context) {
	//TODO

	// a.accountService.Update(ctx,account.AccountUriRequest)
}

func (a *AccountHandler) Delete(ctx *gin.Context) {
	// id := ctx.Param("id")
	// a.accountService.Delete(ctx,account.AccountUriRequest)
}
