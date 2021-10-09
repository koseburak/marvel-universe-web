package marvel

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/koseburak/marvel-universe-web/config"
	"github.com/koseburak/marvel-universe-web/model"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type MarvelClient struct {
	Config     *config.Configuration
	HTTPClient HTTPClient
}

func NewMarvelClient(conf *config.Configuration, httpClient HTTPClient) *MarvelClient {

	return &MarvelClient{
		Config:     conf,
		HTTPClient: httpClient,
	}
}

func (m MarvelClient) GetCharacters(character string) (*model.MarvelResponse, error) {

	// Prepare Marvel Api request headers
	ts := strconv.FormatInt(time.Now().UnixMilli(), 10)
	md5Byte := md5.Sum([]byte(ts + m.Config.MarvelPrivateKey + m.Config.MarvelPublicKey))
	md5Str := fmt.Sprintf("%x", md5Byte)

	baseURL, err := url.Parse(m.Config.MarvelAPIBaseURL)
	if err != nil {
		log.Println("Malformed URL: ", err)
	}

	baseURL.Path += "characters"
	params := url.Values{}
	params.Add("nameStartsWith", character)
	params.Add("hash", md5Str)
	params.Add("apikey", m.Config.MarvelPublicKey)
	params.Add("ts", ts)
	baseURL.RawQuery = params.Encode()

	request, _ := http.NewRequest(http.MethodGet, baseURL.String(), nil)
	response, err := m.HTTPClient.Do(request)
	if err != nil {
		log.Println("Got error while http request to marvel api: ", err)
		return nil, err
	}

	defer response.Body.Close()

	var marvelResponse model.MarvelResponse
	err = json.NewDecoder(response.Body).Decode(&marvelResponse)
	if err != nil {
		log.Println("Got error while unmarshalling response data: ", err)
		return nil, err
	}

	return &marvelResponse, nil
}
