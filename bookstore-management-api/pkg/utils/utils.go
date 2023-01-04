package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ParseBody(req *http.Request, res interface{}) {
	if body, err := ioutil.ReadAll(req.Body); err == nil {
		if err := json.Unmarshal([]byte(body), res); err != nil {
			return
		}
	}
}

func ParseId(req *http.Request) int64 {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	x, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	return x
}
