package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"
)

func main() {

	longURL := input("Digite o URL:")

	status, err := pingerson(longURL)

	if err != nil {
		fmt.Println("url ta zuado")
	}

	fmt.Printf("Url ta funfando? %v \n", status)

	fmt.Println("Encurtando URL...")
	result := ShorterURL(longURL)
	fmt.Printf("%v", result)

}

func ShorterURL(longURL string) string {
	hash := md5.Sum([]byte(longURL))
	hashStr := hex.EncodeToString(hash[:])

	shortURL := hashStr[:8]
	fmtURL := "https://placeholder_da_silva.com/" + shortURL

	shortURLConv := string(fmtURL)

	return shortURLConv
}

func input(prompt string) string {
	fmt.Println(prompt)
	var data string
	fmt.Scanln(&data)
	return data
}

func pingerson(longURL string) (bool, error) {
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	response, err := client.Head(longURL)

	if err != nil {
		return false, err
	}

	defer response.Body.Close()
	return response.StatusCode < 400, nil
}
