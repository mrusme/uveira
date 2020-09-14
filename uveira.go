package main

import (
  "fmt"
  "os"
  "strings"

  "github.com/eliben/gosax"
)


func main() {
  var currentPage Page
  counter := 0
  inPage := false
  inPageTitle :=false

  scb := gosax.SaxCallbacks{
    StartElement: func(name string, attrs []string) {
      switch element := name; element {
        case "page":
          inPage = true
          currentPage = Page{}
          counter++
        case "title":
          inPageTitle = true
        default:
          // TODO
      }
    },

    EndElement: func(name string) {
      switch element := name; element {
        case "page":
          inPage = false
        case "title":
          inPageTitle = false
        default:
          // TODO
      }
    },

    Characters: func(contents string) {
      if inPageTitle {
        currentPage.title = contents
        if len(os.Args) == 3 && strings.Contains(strings.ToLower(contents), strings.ToLower(os.Args[2])) {
          fmt.Printf("%s\n", currentPage.title)
        }
      }
    },
  }

  err := gosax.ParseFile(os.Args[1], scb)
  if err != nil {
    panic(err)
  }

  fmt.Println("--------------------------------------------------------------------------------")
  fmt.Println("Number of pages:", counter)
}
