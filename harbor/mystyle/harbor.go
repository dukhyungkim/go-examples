package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go-examples/harbor/mystyle/model"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	urlProjects = "/projects"
)

type HarborClient interface {
	ListProjects() error
}

type harborClient struct {
	baseURL string
	token   string
}

type HarborConfig struct {
	URL      string
	Username string
	Password string
}

func NewHarborClient(config *HarborConfig) HarborClient {
	return &harborClient{
		baseURL: config.URL + "/api/v2.0",
		token:   base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.Username, config.Password))),
	}
}

func (c *harborClient) ListProjects() error {
	authHeader := fmt.Sprintf("Basic %s", c.token)
	const projectsURI = urlProjects + "?page=1&page_size=10&with_detail=true"
	req, err := http.NewRequest("GET", c.baseURL+projectsURI, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", authHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println(resp.StatusCode)

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var projects []*model.Project
	if err := json.Unmarshal(respData, &projects); err != nil {
		panic(err)
	}
	for _, project := range projects {
		log.Printf("%+v\n", project)
	}
	return nil
}
