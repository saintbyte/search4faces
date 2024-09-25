/*
Module for access to API search4faces.com

Original documention by API: https://search4faces.com/en/api.html
*/
package search4faces

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
)

type API struct {
	APIKey string
	APIURL string
}

func NewSearch4FacesApi() (*API, error) {
	return &API{
		APIKey: "",
		APIURL: ApiURL,
	}, nil
}
func (api *API) RateLimit() error {
	//rateLimit
	RequestObj, err := api.getDefaultRequestData("rateLimit")
	if err != nil {
		return err
	}
	RequestObj.Params = map[string]string{}
	jsonRequest, err := json.Marshal(&RequestObj)
	client, request, _ := api.getRequest(bytes.NewReader(jsonRequest))
	log.Println(client)
	log.Println(request)
	response, e := client.Do(request)
	if e != nil {
		log.Fatal(e)
	}
	log.Println(response)
	return nil
}
func (api *API) DetectFaces() {

}
func (api *API) getDefaultRequestData(method string) (JSONRPCRequst, error) {
	id := uuid.New()
	return JSONRPCRequst{
		Jsonrpc: "2.0",
		Method:  method,
		Id:      id.String(),
	}, nil
}
func (api *API) getApiKey() string {
	value, exists := os.LookupEnv(Search4FacesApiApiKeyEnv)
	if exists {
		return value
	}
	if api.APIKey != "" {
		return api.APIKey
	}
	return ""
}

func (api *API) getRequest(data io.Reader) (*http.Client, *http.Request, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	request, err := http.NewRequest("POST", api.APIURL, data)
	if err != nil {
		return nil, nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "HTTPie/0.9.6")
	request.Header.Set("x-authorization-token", api.getApiKey())
	client := &http.Client{}
	return client, request, nil
}

func (api *API) toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (api *API) DetectFacesFiles(filename string) error {
	// Найти фейсы на картинке , картинка из файла.
	imagesBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	RequestObj, err := api.getDefaultRequestData("detectFaces")
	if err != nil {
		return err
	}
	RequestObj.Params = DetectRequestParams{
		Image: api.toBase64(imagesBytes),
	}
	jsonRequest, err := json.Marshal(&RequestObj)
	client, request, _ := api.getRequest(bytes.NewReader(jsonRequest))
	log.Println(client)
	log.Println(request)
	response, e := client.Do(request)
	if e != nil {
		log.Fatal(e)
	}
	//if response.StatusCode != http.StatusOK {
	//	return errors.New("Http error:" + strconv.Itoa(response.StatusCode))
	//}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	log.Println(string(body))
	return nil
	var result JSONRPCResponse
	err = json.Unmarshal(body, &result)
	log.Println(result)
	//return result.Data[0].Embedding, nil
	return nil
}

func (api *API) SearchFace() {

}
