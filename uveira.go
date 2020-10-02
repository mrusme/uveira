package main

import (
  "fmt"
  // "strings"
  "os"
  "context"
  "time"
  "log"
  "flag"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
)


func main() {
  var queryTitleExact string
  var queryTitleRegEx string

  rtcfg, err := NewRTCFG()
  if err != nil {
    panic(err)
  }

  flag.StringVar(&queryTitleExact, "te", "", "Query by title (exact)")
  flag.StringVar(&queryTitleRegEx, "tr", "", "Query by title (RegEx)")
  flag.Parse()

  client, err := mongo.NewClient(options.Client().ApplyURI(rtcfg.MongoURI))

  ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
  defer cancel()
  err = client.Connect(ctx)

  defer func() {
    if err = client.Disconnect(ctx); err != nil {
        panic(err)
    }
  }()

  collection := client.Database(rtcfg.Database).Collection(rtcfg.Collection)

  var filter bson.M

  if queryTitleExact == "" && queryTitleRegEx != "" {
    filter = bson.M{
      "_id": primitive.Regex{
        Pattern: queryTitleRegEx,
        Options: "i",
      },
    }
  } else if queryTitleExact != "" && queryTitleRegEx == "" {
    filter = bson.M{"_id": queryTitleExact}
  } else {
    log.Fatal("Specify either -te or -tr!")
    os.Exit(-1)
  }

  cur, err := collection.Find(ctx, filter)
  if err != nil {
    log.Fatal(err)
  }

  defer cur.Close(ctx)
  var pages []Page

  for cur.Next(ctx) {
    var page Page

    err := cur.Decode(&page)
    if err != nil {
      log.Fatal(err)
    }

    pages = append(pages, page)
  }

  if err := cur.Err(); err != nil {
    log.Fatal(err)
  }

  foundPages := len(pages)

  if foundPages > 1 {
    // TODO
    log.Println("Multiple results found.")
  } else if foundPages == 1 {
    fmt.Printf("%+v", pages[0].RenderPage())
  } else {
    log.Println("Nothing found.")
  }

  os.Exit(0)
}
