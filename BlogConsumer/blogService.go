package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type BlogConsumerService struct {
	httpClient http.Client
}

func NewBlogConsumerService() *BlogConsumerService {
	return &BlogConsumerService{
		httpClient: http.Client{},
	}
}

func (bc *BlogConsumerService) WriteData(data interface{}, callback string) error {
	jsonbody, err := json.Marshal(data)
	if err != nil {
		return err
	}
	var blog Blog
	if err := json.Unmarshal(jsonbody, &blog); err != nil {
		return err
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(blog)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, callback+"/blog/_doc", &buf)
	req.Header.Add("Content-Type", "application/json")
	resp, err := bc.httpClient.Do(req)
	if err != nil {
		return err
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var esData interface{}
	err = json.Unmarshal(bytes, &esData)
	if err != nil {
		return err
	}
	log.Println("================= Response from es ================")
	log.Println(esData)
	log.Println("===================================================")
	return nil
}

func (bc *BlogConsumerService) DeleteData(data interface{}, callback string) error {
	jsonbody, err := json.Marshal(data)
	if err != nil {
		return err
	}
	var deleteBody DeleteBody
	if err := json.Unmarshal(jsonbody, &deleteBody); err != nil {
		return err
	}
	searchParameter := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"title": deleteBody.Name,
			},
		},
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(searchParameter)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, callback+"/blog/_delete_by_query", &buf)
	req.Header.Add("Content-Type", "application/json")
	resp, err := bc.httpClient.Do(req)
	if err != nil {
		return err
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var esData interface{}
	err = json.Unmarshal(bytes, &esData)
	if err != nil {
		return err
	}
	log.Println("================= Response from es ================")
	log.Println(esData)
	log.Println("===================================================")
	return nil

}
