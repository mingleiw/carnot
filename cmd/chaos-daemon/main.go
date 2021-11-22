package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"entropie.ai/pkg/capture"
)

type example struct {
	test string
}

func test(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))

	c := &capture.Capture{}
	go c.WithIface("lo").WithPort("3000").Start()
}

func main() {
	http.HandleFunc("/test", test)
	log.Println("starting server:")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
