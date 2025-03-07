package infrastructure

import (
	"apiConsumer/src/reservation/application"
	"apiConsumer/src/reservation/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateReservationController struct {
	useCase *application.UpdateReservationUseCase
}

func NewUpdateReservationController(useCase *application.UpdateReservationUseCase) *UpdateReservationController {
	return &UpdateReservationController{useCase: useCase}
}

func (controller *UpdateReservationController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id de order no encontrada"})
		return
	}

	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller.useCase.Run(int32(id), order)

	c.JSON(http.StatusOK, gin.H{
		"message": "Order actualizado exitosamente",
		"data": order})
}
