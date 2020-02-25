package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"

	"go.uber.org/automaxprocs/maxprocs"
)

func main() {
	if os.Getenv("AUTOMAXPROC") == "true" {
		if _, err := maxprocs.Set(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("AUTOMAXPROC enabled")
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type response struct {
	Duration string   `json:"duration"`
	Prime    *big.Int `json:"prime"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	prime := big.NewInt(int64(time.Now().Nanosecond()))

	for i := 0; i < 500; i++ {
		nbr, err := rand.Prime(rand.Reader, 64)
		if err != nil {
			continue
		}
		prime = prime.Add(prime, nbr)
	}

	duration := time.Since(start)
	json.NewEncoder(w).Encode(response{duration.String(), prime})
}
