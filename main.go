package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/btcsuite/btcd/rpcclient"
)

func main() {
	connCfg := &rpcclient.ConnConfig{
		Host:                 "localhost:6662",
		Endpoint:             "/emc",
		User:                 "emcrpcuser",
		Pass:                 "qoe0sT3GHvHCMX8YkY6P87w8EV",
		CookiePath:           "",
		Params:               "",
		DisableTLS:           true,
		Certificates:         []byte{},
		Proxy:                "",
		ProxyUser:            "",
		ProxyPass:            "",
		DisableAutoReconnect: false,
		DisableConnectOnNew:  false,
		HTTPPostMode:         true,
		ExtraHeaders:         map[string]string{},
		EnableBCInfoHacks:    true,
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
