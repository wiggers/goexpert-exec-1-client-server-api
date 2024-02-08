package usecase

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/wiggers/goexpert/exec/1-client-server-api/internal/entity"
)

type checkExchangeRate struct {
	Repository entity.RepositoryInterface
}

func NewCheckExchangeRate(Repository entity.RepositoryInterface) *checkExchangeRate {

	return &checkExchangeRate{
		Repository: Repository,
	}
}

type TypeExchange struct {
	Data entity.Exchange `json:"USDBRL"`
}

type ExchangeOutputDto struct {
	Dollar string
}

func (c *checkExchangeRate) Execute() (*ExchangeOutputDto, error) {

	ctxHttp := context.Background()
	ctxHttp, cancelHttpCtx := context.WithTimeout(ctxHttp, 200*time.Millisecond)
	defer cancelHttpCtx()

	data, err := c.getExchangeData(ctxHttp)
	if err != nil {
		return nil, err
	}

	ctxBD := context.Background()
	ctxBD, cancelCtx := context.WithTimeout(ctxBD, 10*time.Millisecond)
	defer cancelCtx()

	c.Repository.Save(ctxBD, data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	outuput := ExchangeOutputDto{Dollar: data.Bid}
	return &outuput, nil
}

func (c *checkExchangeRate) getExchangeData(ctx context.Context) (*entity.Exchange, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := TypeExchange{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data.Data, nil
}
