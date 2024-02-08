package controller

import (
	"encoding/json"
	"net/http"

	"github.com/wiggers/goexpert/exec/1-client-server-api/internal/entity"
	"github.com/wiggers/goexpert/exec/1-client-server-api/internal/infra/database"
	"github.com/wiggers/goexpert/exec/1-client-server-api/internal/usecase"
	"gorm.io/gorm"
)

type ExchangeController struct {
	Repository entity.RepositoryInterface
}

func NewExchangeController(db *gorm.DB) *ExchangeController {

	repository := database.NewDbRepository(db)
	return &ExchangeController{
		Repository: repository,
	}
}

func (c *ExchangeController) CheckExchangeRate(w http.ResponseWriter, r *http.Request) {

	useCase := usecase.NewCheckExchangeRate(c.Repository)
	response, err := useCase.Execute()

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Bad Request")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
