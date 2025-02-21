package main

import (
	"awesomeProject/Coincap"
	"fmt"
	"log"
	"time"
)

func main() {
	coincapClient, err := Coincap.NewClient(time.Second * 2)
	if err != nil {
		log.Fatal(err)
	}
	assets, err := coincapClient.GetAssets()
	if err != nil {
		log.Fatal(err)
	}
	for _, asset := range assets {
		fmt.Println(asset.Info())
	}
	bitcoin, err := coincapClient.GetAsset("bitcoin")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bitcoin.Info())
}
