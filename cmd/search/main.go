package main

import (
	"github.com/joho/godotenv"
	"github.com/saintbyte/search4faces"
	"os"
)

func main() {
	godotenv.Load()
	_, ok := os.LookupEnv("SEARCH4DACE_API_KEY")
	if !ok {
		panic("Have not key")
	}
	api, err := search4faces.NewSearch4FacesApi()
	if err != nil {
		panic(err)
	}
	err = api.DetectFacesFiles("Richard_Stallman_at_LibrePlanet_2019.jpg")
	if err != nil {
		panic(err)
	}
}
