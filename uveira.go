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
  var queryListTitlesOnly bool

  rtcfg, err := NewRTCFG()
  if err != nil {
    panic(err)
  }

  flag.StringVar(&queryTitleExact, "te", "", "Query by title (exact)")
  flag.StringVar(&queryTitleRegEx, "tr", "", "Query by title (RegEx)")
  flag.BoolVar(&queryListTitlesOnly, "to", false, "List titles only in query results")
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

  var opts *options.FindOptions
  if queryListTitlesOnly == true {
    opts = options.Find().SetProjection(bson.M{"title": 1})
  }
  cur, err := collection.Find(ctx, filter, opts)
  if err != nil {
    log.Fatal(err)
  }

  defer cur.Close(ctx)
  firstPage := true

  for cur.Next(ctx) {
    var page Page

    err := cur.Decode(&page)
    if err != nil {
      log.Fatal(err)
    }

    if firstPage == false && queryListTitlesOnly == false {
      fmt.Printf("---\n\n\n")
    } else if firstPage == true {
      firstPage = false
    }

    fmt.Printf("%+v", page.RenderPage(queryListTitlesOnly))
  }

  if err := cur.Err(); err != nil {
    log.Fatal(err)
  }

  os.Exit(0)
}
