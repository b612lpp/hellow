package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type incValues struct {
	CurrentPercent int `json:"value"`
	CurrentCost    int `json:"cost"`
}
type toFill struct {
	LitersToFill int `json:"liters"`
	CostToFill   int `json:"cost"`
}

func Calc(w http.ResponseWriter, r *http.Request) {
	var v incValues
	q := json.NewDecoder(r.Body)
	q.Decode(&v)
	if ifReqValid(r) && v.CurrentPercent > 0 && v.CurrentCost > 0 {
		fmt.Fprintln(w, doMath(v))
		json.NewEncoder(w).Encode(doMath(v))

	} else {
		fmt.Fprintln(w, "Wrong request")
	}

}

func doMath(a incValues) (b toFill) {
	b.LitersToFill = ((85 - a.CurrentPercent) * 1500) / 100
	b.CostToFill = b.LitersToFill * a.CurrentCost
	return b
}

func ifReqValid(r *http.Request) bool {

	if r.Method == "GET" && r.URL.Path == "/calc/" {
		return true
	}
	return false
}
