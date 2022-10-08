package dto

type CreateCarDto struct {
	Brand string `json:"brand" validate:"required,min=3,max=50"`
	Model string `json:"model" validate:"required,min=3,max=50"`
	Year  int    `json:"year" validate:"required"`
	Color string `json:"color" validate:"required,min=3,max=50"`
}
