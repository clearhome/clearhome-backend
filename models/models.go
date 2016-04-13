package models

import (
  "time"

  "gopkg.in/mgo.v2/bson"
)

type (
  User struct {
    Id            bson.ObjectId   `bson:"_id,omitempty" json:"id"`
    FirstName     string          `json:"firstname"`
    LastName      string          `json:"lastname"`
    Email         string          `json:"email"`
    Password      string          `json:"password,omitempty"`
    HashPassword  []byte          `json:"hashpassword,omitempty"`
  }

  Item struct {
    Id            bson.ObjectId   `bson:"_id,omitempty" json:"id"`
    Owner         string          `json:"owner"`
    Name          string          `json:"name"`
    Description   string          `json:"description"`
    CreatedOn     time.Time       `json:"createdon,omitempty"`
    Private       string          `json:"private,omitonempty"`
    Tags          []string        `json:"tags,omitempty"`
  }
)
