package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

var APIKEY string

func init() {
	godotenv.Load(".env")
	APIKEY = os.Getenv("OPENWEATHER_APIKEY")
}

// openweather rest api (json response encoder struct type)
type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name    string `json:"name"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

// Without Concurrent Example:

func fetchWeather1(city string) WeatherData {
	var data WeatherData

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, APIKEY)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching weather for %s: %s\n", city, err)
		return data
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding weather data for %s: %s\n", city, err)
		return data
	}

	return data
}

func justFetchData() {
	startNow := time.Now()

	cities := []string{
		"London", "Tokyo", "Almaty",
		"Paris", "Toronto", "Seoul",
		"Moscow", "Copenhagen", "Washington",
	}

	for _, city := range cities {
		data := fetchWeather1(city)
		fmt.Println("This is the data", data)
	}

	fmt.Println("This operation took: ", time.Since(startNow))
}

// Concurrent Example:

func fetchWeather2(city string, ch chan<- string, wg *sync.WaitGroup) interface{} {
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	defer wg.Done()

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, APIKEY)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching weather for %s: %s\n", city, err)
		return data
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding weather data for %s: %s\n", city, err)
		return data
	}

	ch <- fmt.Sprintf("This is the %s", city)

	return data
}

func concurrentFetchData() {
	startNow := time.Now()

	cities := []string{
		"London", "Tokyo", "Almaty",
		"Paris", "Toronto", "Seoul",
		"Moscow", "Copenhagen", "Washington",

		"London", "Tokyo", "Almaty",
		"Paris", "Toronto", "Seoul",
		"Moscow", "Copenhagen", "Washington",

		"London", "Tokyo", "Almaty",
		"Paris", "Toronto", "Seoul",
		"Moscow", "Copenhagen", "Washington",
	}

	ch := make(chan string)
	var wg sync.WaitGroup

	for _, city := range cities {
		// add 1 to WaitGroup
		// counter
		wg.Add(1)
		go fetchWeather2(city, ch, &wg)
	}

	// anonymous function
	// wait until goroutines to finish
	// here we have 9 goroutines (cities slice)
	go func() {
		wg.Wait()
		close(ch)
	}()

	// print result of ch, channel
	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("This operation took: ", time.Since(startNow))
}

func main() {
	justFetchData()
	fmt.Println("---")
	concurrentFetchData()
}
