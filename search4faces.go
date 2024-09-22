/*
Module for access to API search4faces.com

Original documention by API: https://search4faces.com/en/api.html
*/
package search4faces

type API struct {
	APIKey string
	APIURL string
}

func NewSearch4FacesApi(ApiKey string) (*API, error) {
	return &API{
		APIKey: ApiKey,
		APIURL: ApiURL,
	}, nil
}
func (api *API) RateLimit() {

}
func (api *API) DetectFaces() {

}
func (api *API) SearchFace() {

}
