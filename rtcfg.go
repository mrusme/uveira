package main

import (
  "os"
  "fmt"
)

// RTCFG is the runtime configuration for Uveira
type RTCFG struct {
  MongoURI string
  Database string
  Collection string
}

// NewRTCFG creates a new runtime configuration
func NewRTCFG() (RTCFG, error) {
  rtcfg := RTCFG{}

  rtcfg.MongoURI = os.Getenv("UVEIRA_MONGO_URI")
  if rtcfg.MongoURI == "" {
    return rtcfg, fmt.Errorf("please specify UVEIRA_MONGO_URI with the mongodb:// connection string")
  }

  rtcfg.Database = os.Getenv("UVEIRA_DATABASE")
  if rtcfg.Database == "" {
    return rtcfg, fmt.Errorf("please specify UVEIRA_DATABASE with the database name (e.g. enwiki)")
  }

  rtcfg.Collection = os.Getenv("UVEIRA_COLLECTION")
  if rtcfg.Collection == "" {
    return rtcfg, fmt.Errorf("please specify UVEIRA_COLLECTION with the collection name (usually pages)")
  }

  return rtcfg, nil
}
