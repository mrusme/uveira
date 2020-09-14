package main

type PageRevision struct {
  id            int
  parentid      int
  timestamp     string
  // TODO: contributor sub-struct
  comment       string
  model         string
  format        string
  text          string
  sha1          string
}

type Page struct {
  title         string
  ns            int
  id            int
  redirectTitle string
  revision      PageRevision
}
