package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type incValues struct {
	Name int `json:"name"`
}

func Calc(w http.ResponseWriter, r *http.Request) {
	var v incValues
	q := json.NewDecoder(r.Body)
	q.Decode(&v)
	fmt.Fprintln(w, doMath(v.Name))

}

func doMath(a int) int {
	return a + 10
}
