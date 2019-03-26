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
		response *http.Response
	)

	//curl -X GET --header 'Accept: application/json' --header 'Authorization: ' --header 'x-date: Mon, 25 Mar 2019 15:43:56 GMT' --header 'Accept-Encoding: gzip' --compressed  'https://ptx.transportdata.tw/MOTC/v2/Rail/THSR/Station?$top=30&$format=JSON'
	client = &http.Client{}
	if request, err = http.NewRequest("GET", c.URL, nil); err != nil {
		log.Printf("failed to link to %v, error is %v", c.URL, err)
		return nil, err
	}
	request.Header.Set("Authorization",`Authorization: hmac username="FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF", algorithm="hmac-sha1", headers="x-date", signature="XuXMaN0J7XrLHX+eP9xdmzWJAyE="`)
	request.Header.Set("x-date", "Tue, 26 Mar 2019 17:00:55 GMT")


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
