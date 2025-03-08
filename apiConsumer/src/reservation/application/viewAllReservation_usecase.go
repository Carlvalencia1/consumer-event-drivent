package application

import "apiConsumer/src/reservation/domain"

type ViewAllReservationUseCase struct {
	mysqlRepository domain.IOrderMysq
}

func NewViewAllReservationUseCase(mysqlRepository domain.IOrderMysq) *ViewAllReservationUseCase {
	return &ViewAllReservationUseCase{mysqlRepository: mysqlRepository}
}

func (uc *ViewAllReservationUseCase) Run() ([]domain.Reservation, error) {
	return uc.mysqlRepository.GetAll()
}
