package worker

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gafernandez/go-hexagonal/internal/core/domain"
	"github.com/gafernandez/go-hexagonal/internal/core/ports"
	"github.com/go-co-op/gocron"
)

type worker struct {
	defRepository ports.DefinanceRepository
}

func NewAssetWorker(repo ports.DefinanceRepository) *worker {
	return &worker{
		defRepository: repo,
	}
}

func (w *worker) Start() {
	w.start60ms()
}

func (w *worker) start60ms() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(60).Seconds().Do(func() {
		assets, err := w.defRepository.GetAllAsset()
		if err != nil {
			//TODO: Handlering error
		}
		for _, e := range assets {
			if e.Worker == "60s" {
				w.doTask(e)
			}
		}
	})
	s.StartBlocking()
}

func (w *worker) doTask(asset domain.Asset) {
	//TODO: URL Factory
	url := "http://localhost:8080/asset/" + asset.Symbol + "/refresh"
	var body interface{}
	_, err := post(url, body, 10)
	if err != nil {
		fmt.Println("ERROR ", err)
		//TODO: Handlering error
	}
}

func post(url string, body interface{}, seconds int64) ([]byte, error) {
	c := http.Client{Timeout: time.Duration(seconds) * time.Second}
	json_data, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	resp, err := c.Post(url, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Invalid http code " + string(resp.StatusCode))
	}

	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error reading body response " + err.Error())
	}

	return bodyResp, nil
}
