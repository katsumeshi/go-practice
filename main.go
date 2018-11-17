package main

import (
    "html/template"
    "log"
    "net/http"
    "encoding/json"
    "time"
    "strconv"
    "fmt"
    "io/ioutil"
)

func main() {
    http.HandleFunc("/number", numberHandler)
    http.HandleFunc("/numberUpdate", numberUpdate)
    http.HandleFunc("/helloWorld", mainHandler)
    http.HandleFunc("/clock", clockHandler)
    http.HandleFunc("/api/helloWorld", jsonHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

type Data struct {
    Number int64
}

var store = Data{0}

func numberHandler(w http.ResponseWriter, r *http.Request) {
    data := store
    t := template.Must(template.ParseFiles("./templates/number.html"))
    t.Execute(w, data)
}

func numberUpdate(w http.ResponseWriter, r *http.Request) {
        b, err := ioutil.ReadAll(r.Body);
        w.Write(b)
        
        r.ParseForm()
        fmt.Println(r.ParseForm())
        calcNumber, err := strconv.ParseInt(r.FormValue("number")[0:], 10, 64)
        if err != nil {
            fmt.Println(err)
        }
        store.Number += calcNumber
        w.Write([]byte(strconv.FormatInt(store.Number, 10)))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles("./templates/helloWorld.html"))
    t.Execute(w, nil)
}

func clockHandler(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles("./templates/clock.html"))
    m := map[string]string{
        "DateTime": time.Now().Format("2006-01-02 15:04:05"),
    }
    t.Execute(w, m)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
    type ResponseBody struct {
        Text string `json:"text"`
    }
    rb := &ResponseBody{
        Text: "Hello World!",
    }
    w.Header().Set("Content-type", "application/json")
    if err := json.NewEncoder(w).Encode(rb); err != nil {
        log.Fatal(err)
    }
}