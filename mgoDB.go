package main

import (
  "fmt"
  "os"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Person struct {
        Name string
        Phone string
}

// connect to the specified mgo bd
func ConnectDB(url string) *mgo.Session{ 
  session, err := mgo.Dial(url)
  if err != nil {
    fmt.Printf("%s",err)
    os.Exit(1)
  }
  session.SetMode(mgo.Monotonic,true)
  return session
}

// insert entries into db
func InsertStruct(session *mgo.Session, DBName string, CollectionName string, object interface{}) {
  if CollectionName == "Calendar" {
      entity,ok := object.(CombinedCalendarItem)
      if ok != true {
      fmt.Printf("%s",ok)
      os.Exit(1)
    }
    fmt.Println(entity)
    c := session.DB(DBName).C(CollectionName)
    err := c.Insert(&entity)
    if err != nil {
        fmt.Printf("%s",err)
        os.Exit(1)
    }
  } else if CollectionName == "Piazza" {
    
  }
}

// get all the information from the db
func GetCalendarResult(session *mgo.Session, DBName string, CollectionName string) []CombinedCalendarItem {
  var results []CombinedCalendarItem
  c := session.DB(DBName).C(CollectionName)
  c.Find(bson.M{}).All(&results)
  return results
}
