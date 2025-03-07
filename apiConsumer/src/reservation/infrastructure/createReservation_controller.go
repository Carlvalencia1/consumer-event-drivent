package infrastructure

import (
	"apiConsumer/src/reservation/application"
	"apiConsumer/src/reservation/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateReservationController struct {
	useCase *application.CreateReservationUseCase
}

func NewCreateOrderController(useCase *application.CreateReservationUseCase) *CreateReservationController {
	return &CreateReservationController{useCase: useCase}
}

func (controller *CreateReservationController) Execute(c *gin.Context) {
	var order domain.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.useCase.Run(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Order creado correctamente",
		"data":   order,
	})
}
