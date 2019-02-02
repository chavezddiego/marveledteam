package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

const privateKey = "a4189f399360fed42ad293d573ef6dc3f7289051"
const publicKey = "16c0baa6064ff23bd850889e90904bbb"

func GetMD5Hash2(keyPub string, keyPri string) string {
	hasher := md5.New()
	hasher.Write([]byte(keyPub))
	hasher.Write([]byte(keyPri))
	return hex.EncodeToString(hasher.Sum(nil))
}
func GetMD5Hash3(keyPub string, keyPri string) string {
	hasher := md5.New()
	io.WriteString(hasher, privateKey)
	io.WriteString(hasher, publicKey)
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	/*
		file, err := os.Create("./magic_msg.txt")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
		if _, err := io.WriteString(file, "Go is fun!"); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	*/

	// .......... TIME ............

	//ts := time.Now()
	//fmt.Println(ts.String())
	/*
		// .......... HASH ..............
		h := md5.New()
		typeh
		fmt.Println(h)
		io.WriteString(h, privateKey)
		io.WriteString(h, publicKey)
		fmt.Printf("%x", h.Sum(nil))
	*/
	/*
		diego := "Hola"

		URL := "http://gateway.marvel.com/v1/public/characters?ts=" + diego + "&apikey=" + publicKey + "&hash=" + diego
		fmt.Println(URL)

		//.... Peticion ......

		resp, err := http.Get(URL) // Realiza la peticion
		if err != nil {
			// handle error
		}
		defer resp.Body.Close() // Cierra la peticion

		body, err := ioutil.ReadAll(resp.Body)
		//*/ // ...
	mensaje := GetMD5Hash2(publicKey, privateKey)
	mensaje2 := GetMD5Hash2(publicKey, privateKey)

	fmt.Println(mensaje)
	fmt.Println(mensaje2)

}
