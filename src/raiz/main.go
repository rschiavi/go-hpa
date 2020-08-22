package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func loopRaiz(x float64, n int) float64 {
	res := 0.0
	for i := 0; i <= n; i++ {
		res += math.Sqrt(x)
	}
	return res
}

func main() {
	resultado := loopRaiz(0.0001, 1000000)
	http.HandleFunc("/", func(res http.ResponseWriter, _ *http.Request) {
		a := fmt.Sprintf("Retorno da função: %f", resultado)
		res.Write([]byte(a))
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
