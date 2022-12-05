package blogdashboardservice

import (
	model "blogging/BlogDashBoard/Model"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Shopify/sarama"
)

var (
	brokersUrl        = []string{"localhost:9092"}
	topic      string = "blogs"
)

type DashBoardService struct {
	httpClient *http.Client
	conn       sarama.SyncProducer
}

func NewDashBoardService() *DashBoardService {
	conn, err := ConnectProducer(brokersUrl)
	if err != nil {
		log.Panicln(err)
	}
	return &DashBoardService{
		httpClient: &http.Client{},
		conn:       conn,
	}
}

func (dbs *DashBoardService) WriteBlog(blog model.Blog) error {
	fmt.Printf("%+v\n", blog)
	requestBody := RequestBody{
		RequestType: "create",
		Body:        blog,
		CallBackUrl: "http://localhost:9200",
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}
	return dbs.PushCommentToQueue(topic, data)
}
func (dbs *DashBoardService) ReadBlog(name string) (interface{}, error) {
	fmt.Println(name)
	searchBody := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"title": name,
			},
		},
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(searchBody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9200/blog/_doc/_search", &buf)
	req.Header.Add("Content-Type", "application/json")
	resp, err := dbs.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var esData interface{}
	err = json.Unmarshal(bytes, &esData)
	if err != nil {
		return nil, err
	}
	return esData, nil
}
func (dbs *DashBoardService) DeleteBlog(name string) error {
	fmt.Println(name)
	requestBody := RequestBody{
		RequestType: "delete",
		Body: DeleteBody{
			Name: name,
		},
		CallBackUrl: "http://localhost:9200",
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}
	return dbs.PushCommentToQueue(topic, data)
}
