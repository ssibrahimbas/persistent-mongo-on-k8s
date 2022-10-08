package mapper

import (
	"time"

	"ssibrahimbas.com/persistent-mongo/car/src/dto"
	"ssibrahimbas.com/persistent-mongo/car/src/entity"
)

func (m *Mapper) MapCreateCarDtoToCarEntity(d *dto.CreateCarDto) *entity.Car {
	return &entity.Car{
		ID:           "",
		Brand:        d.Brand,
		Model:        d.Model,
		Year:         d.Year,
		Color:        d.Color,
		DateOfCreate: time.Now(),
	}
}
