package core

import (
    "apiConsumer/src/core/middleware"
    "apiConsumer/src/reservation/application"
    "apiConsumer/src/reservation/infrastructure"
    "log"

    "github.com/gin-gonic/gin"
)

func InitRoutes() {
    mysqlConn, err := GetDBPool()
    if err != nil {
        log.Fatalf("Error al obtener la conexión a la base de datos: %v", err)
    }

	rabbitmqCh, err := GetChannel()
	if err != nil {
        log.Fatalf("Error al obtener la conexión a la base de datos: %v", err)
    }

    mysqlRepository := infrastructure.NewMysqlRepository(mysqlConn.DB)
	rabbitqmRepository := infrastructure.NewRabbitRepository(rabbitmqCh.ch)

    createReservationUseCase := application.NewCreateReservationUseCase(rabbitqmRepository, mysqlRepository)
    updateReservationUseCase := application.NewUpdateReservationUseCase(mysqlRepository)
    deleteReservationUseCase := application.NewDeleteReservationUseCase(mysqlRepository)
    getAllReservationUseCase := application.NewViewAllReservationUseCase(mysqlRepository)
    getReservationByIdUseCase := application.NewViewReservationByIdUseCase(mysqlRepository)
    getReservationByCellphoneUseCase := application.NewViewByCellphoneReservationUseCase(mysqlRepository)

    createOrderController := infrastructure.NewCreateOrderController(createReservationUseCase)
    updateReservationController := infrastructure.NewUpdateReservationController(updateReservationUseCase)
    deleteReservationController := infrastructure.NewDeleteReservationController(deleteReservationUseCase)
    getAllReservationController := infrastructure.NewViewAllReservationController(getAllReservationUseCase)
    getReservationByIdController := infrastructure.NewViewByIdReservationController(getReservationByIdUseCase)
    getReservationByCellphoneController := infrastructure.NewViewByCellphoneReservationController(getReservationByCellphoneUseCase)

    router := gin.Default()
    corsMiddleware := middleware.NewCorsMiddleware()
    router.Use(corsMiddleware)

    router.POST("/reservation", createOrderController.Execute)
    router.PUT("/reservation/:id", updateReservationController.Execute)
    router.DELETE("/reservation/:id", deleteReservationController.Execute)
    router.GET("/reservation", getAllReservationController.Execute)
    router.GET("/reservation/:id", getReservationByIdController.Execute)
    router.GET("/reservations/cellphone/:cellphone", getReservationByCellphoneController.Execute)

    if err := router.Run(":8082"); err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }
}
