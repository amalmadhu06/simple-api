package handler

import (
	"fmt"
	"net/http"

	services "github.com/amalmadhu06/simple-api/pkg/usecase/interfaces"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

func (cr *UserHandler) FindAll(c *gin.Context) {
	users, err := cr.userUseCase.FindAll(c.Request.Context())
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, users)
}

type Xyz struct {
	Name string
	Age  int
}

func NewXyz(name string, age int) *Xyz {
	return &Xyz{
		Name: name,
		Age:  age,
	}
}

func abc() {
	x := NewXyz("Amal", 25)
	fmt.Println(x.Name)
	fmt.Println(x.Age)
}
