package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/btcsuite/btcd/rpcclient"
)

func maingpt() {
	connCfg := &rpcclient.ConnConfig{
		Host:         "localhost:6662",
		User:         "emcrpcuser",
		Pass:         "qoe0sT3GHvHCMX8YkY6P87w8EV",
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		blockCount, err := client.GetBlockCount()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Block count: %d", blockCount)
	})

	log.Fatal(http.ListenAndServe(":8880", nil))
}
