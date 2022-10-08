package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/ssi-core/pkg/result"
	"ssibrahimbas.com/persistent-mongo/car/src/dto"
)

func (h *Handler) Create(c *fiber.Ctx) error {
	l, a := h.i18n.GetLanguagesInContext(c)
	d := &dto.CreateCarDto{}
	h.parseBody(c, d)
	h.s.Create(d, l, a)
	return result.Success(h.i18n.Translate("car_created", l, a), fiber.StatusCreated)
}

func (h *Handler) GetList(c *fiber.Ctx) error {
	l, a := h.i18n.GetLanguagesInContext(c)
	cars := h.s.GetList()
	return result.SuccessData(h.i18n.Translate("car_list", l, a), cars, fiber.StatusOK)
}

func (h *Handler) parseBody(c *fiber.Ctx, d interface{}) {
	l, a := h.i18n.GetLanguagesInContext(c)
	if err := c.BodyParser(d); err != nil {
		panic(result.Error(h.i18n.Translate("invalid_request", l, a), fiber.StatusBadRequest))
	}
	h.validateStruct(d, l, a)
}

func (h *Handler) validateStruct(d interface{}, l, a string) {
	if errors := h.v.ValidateStruct(d, l, a); len(errors) > 0 {
		panic(result.ErrorData(h.i18n.Translate("invalid_request", l, a), errors, fiber.StatusBadRequest))
	}
}
