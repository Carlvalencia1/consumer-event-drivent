package domain

type IOrderRabbitqm interface {
	Save(reservation *Reservation) error
}
