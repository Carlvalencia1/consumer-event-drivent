package domain

type IOrderMysq interface {
	Save(order *Reservation) error
	GetById(id int32) (*Reservation, error)
	GetByCellphone(cellphone int32) ([]Reservation, error)
	GetAll() ([]Reservation, error)
	Update(id int32, reservation Reservation) error
	Delete(id int32) error
}
