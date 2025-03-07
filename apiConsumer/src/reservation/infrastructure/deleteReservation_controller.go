package infrastructure

import (
	"apiConsumer/src/reservation/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteReservationController struct {
	useCase *application.DeleteReservationUseCase
}

func NewDeleteReservationController(useCase *application.DeleteReservationUseCase) *DeleteReservationController {
	return &DeleteReservationController{useCase: useCase}
}

func (controller *DeleteReservationController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id de order no encontrada"})
		return
	}

	controller.useCase.Run(int32(id))

	c.JSON(http.StatusOK, gin.H{"estatus": "Orden eliminado correctamente"})
}
