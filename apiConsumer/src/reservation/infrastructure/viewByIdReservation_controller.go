package infrastructure

import (
	"apiConsumer/src/reservation/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ViewByIdReservationController struct {
	useCase *application.ViewByIdReservationUseCase
}

func NewViewByIdReservationController(useCase *application.ViewByIdReservationUseCase) *ViewByIdReservationController {
	return &ViewByIdReservationController{useCase: useCase}
}

func (controller *ViewByIdReservationController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id no encontrada"})
		return
	}

	order, err := controller.useCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orden no encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Order": order})
}
