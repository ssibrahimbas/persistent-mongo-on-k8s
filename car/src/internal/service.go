package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/ssi-core/pkg/helper"
	"github.com/ssibrahimbas/ssi-core/pkg/i18n"
	"github.com/ssibrahimbas/ssi-core/pkg/result"
	"ssibrahimbas.com/persistent-mongo/car/src/dto"
	"ssibrahimbas.com/persistent-mongo/car/src/entity"
	"ssibrahimbas.com/persistent-mongo/car/src/mapper"
)

type Service struct {
	r    *Repo
	i18n *i18n.I18n
	m    *mapper.Mapper
}

type ServiceConfig struct {
	Repo *Repo
	I18n *i18n.I18n
}

func NewService(c *ServiceConfig) *Service {
	return &Service{
		r:    c.Repo,
		i18n: c.I18n,
		m:    mapper.New(),
	}
}

func (s *Service) Create(d *dto.CreateCarDto, l ...string) {
	_, err := s.r.Create(s.m.MapCreateCarDtoToCarEntity(d))
	if err != nil {
		panic(result.Error(err.Error(), fiber.StatusBadRequest))
	}
}

func (s *Service) GetList() []*entity.Car {
	c, err := s.r.GetList()
	helper.CheckErr(err)
	return c
}
