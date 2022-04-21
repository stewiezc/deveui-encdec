package main

import (
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
)

func main() {
	// discover flags
	decode := flag.String("d", "", "string to decode")
	encode := flag.String("e", "", "string to encode")
	flag.Parse()

	if *decode != "" {
		data := decode

		// Base64 Standard Decoding
		sDec, err := base64.StdEncoding.DecodeString(*data)
		if err != nil {
			fmt.Printf("Error decoding string: %s ", err.Error())
			return
		}
		hx := hex.EncodeToString([]byte(sDec))

		fmt.Println(string(hx))
	}

	if *encode != "" {
		data := encode

		// hex Decoding
		hx, err := hex.DecodeString(*data)
		if err != nil {
			fmt.Printf("Error decoding string: %s ", err.Error())
			return
		}
		sEnc := base64.StdEncoding.EncodeToString([]byte(hx))

		fmt.Println(string(sEnc))
	}

}
