package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type prometheusResponse struct {
	Data struct {
		Result []struct {
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

func queryPrometheus(query string) (float64, error) {
	url := "http://localhost:9090/api/v1/query?query=" + query

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var result prometheusResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, err
	}

	if len(result.Data.Result) == 0 {
		return 0, fmt.Errorf("no result for query: %s", query)
	}

	valueString := result.Data.Result[0].Value[1].(string)

	value, err := strconv.ParseFloat(valueString, 64)
	if err != nil {
		return 0, err
	}

	return value, nil
}
