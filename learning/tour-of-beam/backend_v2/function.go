package api

import (
    "io/ioutil"
    "net/http"
)

// FileTest func
func FileTest(w http.ResponseWriter, r *http.Request) {
    content, err := ioutil.ReadFile("./test.txt")
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(err.Error()))
        return
    }

    w.Write(content)
}