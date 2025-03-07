package application

import "apiConsumer/src/reservation/domain"

type DeleteReservationUseCase struct {
	mysqlRepository domain.IOrderMysq
}

func NewDeleteReservationUseCase(mysqlRepository domain.IOrderMysq) *DeleteReservationUseCase {
	return &DeleteReservationUseCase{mysqlRepository: mysqlRepository}
}

func (uc *DeleteReservationUseCase) Run(id int32) error {
	return uc.mysqlRepository.Delete(id)
}
