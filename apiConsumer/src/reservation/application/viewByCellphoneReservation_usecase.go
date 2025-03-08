package application

import "apiConsumer/src/reservation/domain"

type ViewByCellphoneReservationUseCase struct {
	mysqlRepository domain.IOrderMysq
}

func NewViewByCellphoneReservationUseCase(mysqlRepository domain.IOrderMysq) *ViewByCellphoneReservationUseCase {
	return &ViewByCellphoneReservationUseCase{mysqlRepository: mysqlRepository}
}

func (uc *ViewByCellphoneReservationUseCase) Run(cellphone int32) ([]domain.Reservation, error) {
	return uc.mysqlRepository.GetByCellphone(cellphone)
}
