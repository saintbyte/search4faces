package main

import "github.com/saintbyte/search4faces"

func main() {
	api, err := search4faces.NewSearch4FacesApi()
	if err != nil {
		panic(err)
	}

}
