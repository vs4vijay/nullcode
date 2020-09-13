package main

import (
	"fmt"
	"log"
	"net/http"
	"nullcode/pkg/nullcode"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
)

var (
	port = os.Getenv("PORT")
)

func main() {
	fmt.Println("Starting Service")

	svc := nullcode.CodeService{}

	addHandler := httptransport.NewServer(
		nullcode.MakeAddEndpoint(svc),
		nullcode.DecodeAddRequest,
		nullcode.EncodeAddResponse,
	)

	http.Handle("/api/v1/codes", addHandler)

	if port == "" {
		port = "7777"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
