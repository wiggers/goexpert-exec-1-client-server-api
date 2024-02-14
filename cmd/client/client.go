package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Data struct {
	Value string `json:"Dollar"`
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	stop(err)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Timeout exceeded")
		panic("Timeout exceeded")
	}

	body, err := io.ReadAll(resp.Body)
	stop(err)

	data := Data{}
	err = json.Unmarshal(body, &data)
	stop(err)

	f, err := os.Create("cotacao.txt")
	stop(err)

	f.WriteString("Dólar:" + data.Value)
	f.Close()
}

func stop(err error) {
	if err != nil {
		panic(err)
	}
}
