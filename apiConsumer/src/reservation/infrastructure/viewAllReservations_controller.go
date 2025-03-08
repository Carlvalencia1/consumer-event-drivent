package infrastructure

import (
	"apiConsumer/src/reservation/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ViewAllReservationController struct {
	useCase *application.ViewAllReservationUseCase
}

func NewViewAllReservationController(useCase *application.ViewAllReservationUseCase) *ViewAllReservationController {
	return &ViewAllReservationController{useCase: useCase}
}

func (controller *ViewAllReservationController) Execute(c *gin.Context) {
	reservations, err := controller.useCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reservations": reservations})
}
