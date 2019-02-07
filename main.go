package main
 
import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

type response1 struct {
	Version   string
    Lastcommitsha string
	Description string
       
}

type response2 struct {
	Myapplication []string
       
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
   
     res2D := &response1{
        Version:   "1.0",
        Lastcommitsha: "asdfsdfasdf",
        Description: "pre-interview technical test"}
    strA, _ := json.Marshal(res2D)
  
  res3D := &response2{
        Myapplication: []string{(string(strA))}}
    fmt.Println(res3D)

}
func main() {
    http.HandleFunc("/health", ServeHTTP)
    http.ListenAndServe("localhost:8080", nil)
}
