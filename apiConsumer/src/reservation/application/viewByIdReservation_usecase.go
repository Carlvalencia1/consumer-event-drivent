package application

import "apiConsumer/src/reservation/domain"

type ViewByIdReservationUseCase struct {
	mysqlRepository domain.IOrderMysq
}

func NewViewReservationByIdUseCase(mysqlRepository domain.IOrderMysq) *ViewByIdReservationUseCase {
	return &ViewByIdReservationUseCase{mysqlRepository: mysqlRepository}
}

func (uc *ViewByIdReservationUseCase) Run(id int32) (*domain.Reservation, error) {
	return uc.mysqlRepository.GetById(id)
}
