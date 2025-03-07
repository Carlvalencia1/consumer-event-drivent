package application

import "apiConsumer/src/reservation/domain"

type UpdateReservationUseCase struct {
	mysqlRepository domain.IOrderMysq
}

func NewUpdateReservationUseCase(mysqlRepository domain.IOrderMysq) *UpdateReservationUseCase {
	return &UpdateReservationUseCase{mysqlRepository: mysqlRepository}
}

func (uc *UpdateReservationUseCase) Run(id int32, order domain.Order) error {
	return uc.mysqlRepository.Update(id, order)
}
