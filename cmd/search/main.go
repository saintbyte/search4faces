package main

import (
	"github.com/saintbyte/search4faces"
	"os"
)

func main() {
	apiKey, ok := os.LookupEnv("SEARCH4DACE_API_KEY")
	if !ok {
		panic("Have not key")
	}
	api, err := search4faces.NewSearch4FacesApi(apiKey)
	if err != nil {
		panic(err)
	}
	api.DetectFacesFiles("Richard_Stallman_at_LibrePlanet_2019.jpg")
}
