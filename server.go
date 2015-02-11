package main

import(
  "net/http"
  _"encoding/json"
  _"io"
  "io/ioutil"
  "fmt"
  "os"
  _"gopkg.in/mgo.v2"
)

type HelloMessage struct{
  msg string
}

var myMsg HelloMessage

func CalendarHandler(w http.ResponseWriter, req *http.Request){
    
  switch req.Method {
    case "GET":
    response,err := http.Get("https://www.googleapis.com/calendar/v3/calendars/nvqeh6qvdhrhnj0bn6h2gkg2io%40group.calendar.google.com/events?key=AIzaSyDVjq2FfXV-IUFw7vnsBymz_eCG91hddS0") 
      if err != nil {
        fmt.Printf("%s",err)
        os.Exit(1)
      }
      content,err := ioutil.ReadAll(response.Body)
      if err != nil {
        fmt.Printf("%s",err)
        os.Exit(1)
      }
      calendarObj := CalendarParser(content)
      w.Write(calendarObj)
      fmt.Println(req.URL.Path[len("/Calendar/"):])
    
    default:
      w.WriteHeader(400)
  }
}

func main(){
  s := Reverse("hello")
  fmt.Println(s)
  http.HandleFunc("/Calendar/",CalendarHandler)
  http.ListenAndServe(":8080",nil)
}