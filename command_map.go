package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"

	"github.com/BredSandowich/pokedexcli/internal/pokecache"
)

type config struct {
	Next *string
	Previous *string
	cache pokecache.Cache
}

type LocationAreaResult struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type LocationAreaResponse struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []LocationAreaResult `json:"results"`
}



func getMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if cfg.Next != nil {
		url = *cfg.Next
	}
	
	//Check if URL is already in cache
	if val, ok := cfg.cache.Get(url); ok{
		var locationResp LocationAreaResponse
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			log.Fatal(err)
		}
		cfg.Next = locationResp.Next
		cfg.Previous = locationResp.Previous

		for _, place := range locationResp.Results {
			fmt.Println(place.Name)
		}
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode > 299 {
		log.Fatalf("Pokedex api request failed with status code %d", res.StatusCode)
	}
	
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	cfg.cache.Add(url, body)
	

	var locationResp LocationAreaResponse
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		log.Fatal(err)
	}
	cfg.Next = locationResp.Next
	cfg.Previous = locationResp.Previous

	for _, place := range locationResp.Results {
		fmt.Println(place.Name)
	}
	return nil
}


func getMapBack(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("You're on the first page.")
		return nil
	}
	url := *cfg.Previous
	
	//Check if URL is already in cache
	if val, ok := cfg.cache.Get(url); ok{
		var locationResp LocationAreaResponse
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			log.Fatal(err)
		}
		cfg.Next = locationResp.Next
		cfg.Previous = locationResp.Previous

		for _, place := range locationResp.Results {
			fmt.Println(place.Name)
		}
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode > 299 {
		log.Fatalf("Pokedex api request failed with status code %d", res.StatusCode)
	}
	
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	cfg.cache.Add(url, body)

	var locationResp LocationAreaResponse
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		log.Fatal(err)
	}

	cfg.Next = locationResp.Next
	cfg.Previous = locationResp.Previous

	for _, place := range locationResp.Results {
		fmt.Println(place.Name)
	}
	return nil
}
