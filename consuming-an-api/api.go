package main

import (
    "fmt"
    "io/ioutil"
    "os"    
    "log"
    "net/http"
    "encoding/json"     
    "github.com/julienschmidt/httprouter"   
)

type Message struct {
    Body string
    Number int8
    Decimal float32
    Validate bool
}

type Pokemon struct {
    name string
    level int8
}

func returnJson(url string, w http.ResponseWriter, r *http.Request){
    fmt.Print("aqui")
    response, err := http.Get(url)

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
   
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, string(responseData)) 
}

func retornarUsuarioAleatorio(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
    fmt.Print("aqui")
    returnJson("https://randomuser.me/api/", w, r)
}

func retornarPokemon(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
    nomePokemon := ps.ByName("nome")
    urlApi := "https://pokeapi.co/api/v2/pokemon/" + nomePokemon

    returnJson(urlApi, w, r)
}

func retornarStruct(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
    m := Message{"Hello, Mundão!", 124, 1687.87845, true}
    b, err := json.Marshal(m)
    if err != nil {
        log.Fatal(err)
    }

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, string(b)) 
}

func handleRequests() {
    router := httprouter.New()
    router.GET("/retornarUsuarioAleatorio", retornarUsuarioAleatorio)
    router.GET("/retornarStruct", retornarStruct)
    router.GET("/retornarPokemon/:nome", retornarPokemon)
    //router.POST("/criarPokemon") TODO

    log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
    handleRequests()
}
