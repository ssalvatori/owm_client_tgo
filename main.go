package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const openWeatherMapEndpoint = "https://api.openweathermap.org/data/2.5"

func buildURL(apiKey string, cityID string, lang string) string {
	return fmt.Sprintf("%s/weather?id=%s&appid=%s&units=metric&lang=%s", openWeatherMapEndpoint, cityID, apiKey, lang)
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

//OpenWeatherResponde response
type OpenWeatherResponde struct {
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure float64 `json:"pressure"`
		Humidity float64 `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

//Config structure
type Config struct {
	Setup struct {
		APIKey      string `json:"api_key"`
		Lang        string `json:"lang"`
		ArgIndex    int    `json:"arg_index"`
		DefaultCity string `json:"default_city"`
		LookupCity  string
	} `json:"setup"`
	CitiesMap []CitiesMapKeyValue `json:"cities_map"`
}

//CitiesMapKeyValue configuration
type CitiesMapKeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func getConfig(args []string) *Config {
	file, _ := ioutil.ReadFile("config.json")

	config := new(Config)

	err := json.Unmarshal([]byte(file), &config)

	if err != nil {
		fmt.Printf("Couldn't parse config.json")
		os.Exit(1)
	}

	if len(args) < config.Setup.ArgIndex {
		config.Setup.LookupCity = config.Setup.DefaultCity
	} else {
		config.Setup.LookupCity = args[config.Setup.ArgIndex]
	}

	return config
}

//getCity read cities maps and return the lookup city if doesn´t exist it will return the first one or scl
func getCity(cities []CitiesMapKeyValue, lookup string) string {

	if len(cities) > 0 {
		for _, value := range cities {
			if value.Key == lookup {
				return value.Value
			}
		}
		return cities[0].Value
	}
	return "3873544" //SCL cityID
}

func main() {

	config := getConfig(os.Args)

	res := new(OpenWeatherResponde)
	city := getCity(config.CitiesMap, config.Setup.LookupCity)
	getJSON(buildURL(config.Setup.APIKey, city, config.Setup.Lang), res)
	fmt.Printf("Weather for [%s] actual [%2.fºC]  min %2.fºC max %2.fºC\n", res.Name, res.Main.Temp, res.Main.TempMin, res.Main.TempMax)
}
