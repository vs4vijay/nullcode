package main

import (
	"fmt"
	"log"
	"net/http"
	"nullcode/pkg/nullcode"

	httptransport "github.com/go-kit/kit/transport/http"
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

	log.Fatal(http.ListenAndServe(":7777", nil))
}
