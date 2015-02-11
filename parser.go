package main

import(
  "encoding/json"
  "fmt"
  _"gopkg.in/mgo.v2"
)

func CalendarParser( jsonStream []byte ) []byte {
  var c Calendar
  err := json.Unmarshal(jsonStream, &c)
  if err != nil {
    fmt.Printf("error: %s",err)
  }
 //buf,_ := json.Marshal(c)
  
  // insert into db for every entry
  var session = ConnectDB("localhost")
  for i:=0; i< len(c.Items); i++{
   entry :=  CombinedCalendarItem{c.Items[i].Summary, c.Items[i].Description,c.Items[i].Location, c.Items[i].Start.DateTime}
      //fmt.Println(entry)
      InsertStruct(session,"HackRice","Calendar",entry)
  }
  
  // get all the entries from DB
  results := GetCalendarResult(session, "HackRice", "Calendar")
  //fmt.Println(results[0])
  buf,_ := json.Marshal(results)
  
  return buf
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}