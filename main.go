package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
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

	return hex.EncodeToString(hasher.Sum(nil))
}

func imprimir(resp Response) {
	for i := 0; i < len(resp.Data.Results); i++ {
		fmt.Println("Name: " + resp.Data.Results[i].Name)
		fmt.Println(resp.Data.Results[i].ID)
		fmt.Println("Description: " + resp.Data.Results[i].Description)

		for j := 0; j < len(resp.Data.Results[i].Comics.Items); j++ {
			fmt.Println("Comic: " + resp.Data.Results[i].Comics.Items[j].Name)
		}
		for j := 0; j < len(resp.Data.Results[i].Series.Items); j++ {
			fmt.Println("Serie: " + resp.Data.Results[i].Series.Items[j].Name)
		}

		fmt.Println("------------------------------------------")
	}
}
func imprimirLista(resp Response) {
	for i := 0; i < len(resp.Data.Results); i++ {
		fmt.Println("Name: " + resp.Data.Results[i].Name)
	}
}

func main() {
	//var opcion int
	var heroe string
	var opcion int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Elegi una opcion: ")
	fmt.Println("1- Busca a tu heroe")
	fmt.Println("2- Lista de los primeros 20 heroes")

	fmt.Scanf("%d", &opcion)
	switch opcion {
	case 1:
		fmt.Print("Elige un heroe: ")
		entrada, _ := reader.ReadString('\n')      // Leer hasta el separador de salto de línea
		heroe = strings.TrimRight(entrada, "\r\n") // Remover el salto de línea de la entrada del usuario

		hash := GetMD5Hash(publicKey, privateKey)
		ts := time.Now().Format("20060102150405")

		hero, err := url.Parse(heroe)
		if err != nil {
			fmt.Println(err)
		}

		URL := "http://gateway.marvel.com/v1/public/characters?name=" + hero.String() + "&ts=" + ts + "&apikey=" + publicKey + "&hash=" + hash
		// ************************
		resp, err := http.Get(URL)

		if err != nil {
			log.Fatal(err)
		}

		responseData, err := ioutil.ReadAll(resp.Body)

		var responseObject Response
		json.Unmarshal(responseData, &responseObject)
		// ********************************

		// Imprimir todos los datos
		imprimir(responseObject)

	case 2:

		hash := GetMD5Hash(publicKey, privateKey)
		ts := time.Now().Format("20060102150405")

		URL := "http://gateway.marvel.com/v1/public/characters?ts=" + ts + "&apikey=" + publicKey + "&hash=" + hash
		// ***********************
		resp, err := http.Get(URL)

		if err != nil {
			log.Fatal(err)
		}

		responseData, err := ioutil.ReadAll(resp.Body)

		var responseObject Response
		json.Unmarshal(responseData, &responseObject)
		// ********************************
		// No entiendo este punto, debo solo imprimir una lista con los nombres o tener todas las estructuras en una array
		// para despues usarlo en otro caso
		// Si es la segunda en responseObject se encuentra todos los datos de los 20 heroes ordenados
		// para ser usado en otros problemas / ejercicios
		// Por ejemplo el id del primer heroe          responseObject.Data.Results[0].ID
		// Imprimir todos los datos
		imprimirLista(responseObject)
	}
}
