package infrastructure

import (
	"apiConsumer/src/reservation/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ViewByCellphoneReservationController struct {
	useCase *application.ViewByCellphoneReservationUseCase
}

func NewViewByCellphoneReservationController(useCase *application.ViewByCellphoneReservationUseCase) *ViewByCellphoneReservationController {
	return &ViewByCellphoneReservationController{useCase: useCase}
}

func (controller *ViewByCellphoneReservationController) Execute(c *gin.Context) {
	cellphoneStr := c.Param("cellphone")
	cellphone, err := strconv.Atoi(cellphoneStr)

	orders, err := controller.useCase.Run(int32(cellphone))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Órdenes no encontradas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}
