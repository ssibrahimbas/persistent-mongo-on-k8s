package internal

import (
	"github.com/ssibrahimbas/ssi-core/pkg/http"
	"github.com/ssibrahimbas/ssi-core/pkg/i18n"
	"github.com/ssibrahimbas/ssi-core/pkg/validator"
	"ssibrahimbas.com/persistent-mongo/car/src/config"
)

type Handler struct {
	s    *Service
	c    *config.App
	h    *http.Client
	i18n *i18n.I18n
	v    *validator.Validator
}

type HandlerConfig struct {
	Service    *Service
	Config     *config.App
	HttpClient *http.Client
	Validator  *validator.Validator
	I18n       *i18n.I18n
}

func NewHandler(c *HandlerConfig) *Handler {
	return &Handler{
		s:    c.Service,
		c:    c.Config,
		h:    c.HttpClient,
		v:    c.Validator,
		i18n: c.I18n,
	}
}

func (h *Handler) InitAllVersions() {
	h.initV1()
}

func (h *Handler) initV1() {
	v1 := h.h.App.Group("api/cars/v1")
	v1.Use(h.i18n.I18nMiddleware)
	v1.Post("/create", h.Create)
	v1.Get("/", h.GetList)
}
