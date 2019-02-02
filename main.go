package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Response struct {
	Data Data `json:"data"`
}
type Data struct {
	Results []Results `json:"results"`
}

type Results struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Comics      Comics `json:"comics"`
	Series      Series `json:"series"`
}

type Comics struct {
	Available     int     `json:"available"`
	CollectionURI string  `json:"collectionURI"`
	Items         []Items `json:"items"`
}
type Items struct {
	Name string `json:"name"`
}
type Series struct {
	Items []Items `json:"items"`
}

const privateKey = "a4189f399360fed42ad293d573ef6dc3f7289051"
const publicKey = "16c0baa6064ff23bd850889e90904bbb"

func GetMD5Hash(keyPub string, keyPri string) string {
	ts := time.Now().Format("20060102150405")
	hasher := md5.New()
	hasher.Write([]byte(ts))
	hasher.Write([]byte(privateKey))
	hasher.Write([]byte(publicKey))

	//io.WriteString(hasher, keyPub)
	//io.WriteString(hasher, keyPri)
	return hex.EncodeToString(hasher.Sum(nil))
}

func getConnection() {

	hash := GetMD5Hash(publicKey, privateKey)
	ts := time.Now().Format("20060102150405")

	// fmt.Printf("%x", hash)

	URL := "http://gateway.marvel.com/v1/public/characters?ts=" + ts + "&apikey=" + publicKey + "&hash=" + hash

	//fmt.Println(URL)
	resp, err := http.Get(URL)

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	// Traer todos los elementos de la API
	/*
		// Imprimir todos los datos
		for i := 0; i < len(responseObject.Data.Results); i++ {
			fmt.Println("Name: " + responseObject.Data.Results[i].Name)
			fmt.Println(responseObject.Data.Results[i].ID)
			fmt.Println("Description: " + responseObject.Data.Results[i].Description)

			for j := 0; j < len(responseObject.Data.Results[i].Comics.Items); j++ {
				fmt.Println("Comic: " + responseObject.Data.Results[i].Comics.Items[j].Name)
			}
			for j := 0; j < len(responseObject.Data.Results[i].Series.Items); j++ {
				fmt.Println("Serie: " + responseObject.Data.Results[i].Series.Items[j].Name)
			}

			fmt.Println("------------------------------------------")
		}
	*/
	// -------------------------------------------------

	//fmt.Println(responseObject.Data.Results[0])
}

func buscarHeroe(heroe string) {
	fmt.Println(heroe)
	hash := GetMD5Hash(publicKey, privateKey)
	ts := time.Now().Format("20060102150405")

	// fmt.Printf("%x", hash)

	hero, err := url.Parse(heroe)
	if err != nil {
		fmt.Println(err)
	}
	heroeposta := hero.String()
	fmt.Println(heroeposta)
	URL := "http://gateway.marvel.com/v1/public/characters?name=" + heroeposta + "?ts=" + ts + "&apikey=" + publicKey + "&hash=" + hash

	fmt.Println(URL)
	/*
		resp, err := http.Get(URL)

		if err != nil {
			log.Fatal(err)
		}

		responseData, err := ioutil.ReadAll(resp.Body)

		var responseObject Response
		json.Unmarshal(responseData, &responseObject)
	*/
	// Traer todos los elementos de la API
	/*
		// Imprimir todos los datos
		for i := 0; i < len(responseObject.Data.Results); i++ {
			fmt.Println("Name: " + responseObject.Data.Results[i].Name)
			fmt.Println(responseObject.Data.Results[i].ID)
			fmt.Println("Description: " + responseObject.Data.Results[i].Description)

			for j := 0; j < len(responseObject.Data.Results[i].Comics.Items); j++ {
				fmt.Println("Comic: " + responseObject.Data.Results[i].Comics.Items[j].Name)
			}
			for j := 0; j < len(responseObject.Data.Results[i].Series.Items); j++ {
				fmt.Println("Serie: " + responseObject.Data.Results[i].Series.Items[j].Name)
			}

			fmt.Println("------------------------------------------")
		}
	*/
	// -------------------------------------------------

	//fmt.Println(responseObject.Data.Results[0])
}

func main() {
	//var opcion int
	var heroe string
	fmt.Printf("Elige un heroe: ")
	fmt.Scanf("%s", &heroe) /*
		fmt.Printf("Elegi una opcion: ")
		fmt.Scanf("%d", &opcion)
		switch opcion {
		case 1:
		case 2:
			fmt.Print("Opcion 2")
		}
	*/
	fmt.Printf("%s", heroe) /*
		hash := GetMD5Hash(publicKey, privateKey)
		ts := time.Now().Format("20060102150405")

		// fmt.Printf("%x", hash)

		hero, err := url.Parse(heroe)
		if err != nil {
			fmt.Println(err)
		}
		heroeposta := hero.String()
		fmt.Println(heroeposta)
		URL := "http://gateway.marvel.com/v1/public/characters?name=" + heroeposta + "?ts=" + ts + "&apikey=" + publicKey + "&hash=" + hash

		fmt.Println(URL)
	*/
	//getConnection()
}
