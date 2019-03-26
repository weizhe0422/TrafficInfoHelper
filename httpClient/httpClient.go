package httpClient

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	URL string
}

func InitHttpClient(URL string) *Client {
	var (
		c *Client
	)
	c = new(Client)
	c.URL = URL
	return c
}

func (c *Client) GetHttpResp() (respBody []byte, err error) {
	var (
		client   *http.Client
		request  *http.Request
		response []http.Response
	)

	client = &http.Client{}
	if request, err = http.NewRequest("GET", c.URL, nil); err != nil {
		log.Printf("failed to link to %v, error is %v", c.URL, err)
		return nil, err
	}
	if response, err = client.Do(request); err != nil {
		log.Printf("failed to get response from  %v, error is %v", c.URL, err)
		return nil, err
	}
	if respBody, err = ioutil.ReadAll(response.Body); err != nil {
		log.Printf("failed to parse response from  %v, error is %v", c.URL, err)
		return nil, err
	}
	defer response.Body.Close()

	return respBody, nil
}
