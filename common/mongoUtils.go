package common

import (
  "gopkg.in/mgo.v2"
)

var session *mgo.Session

func GetSession() *mgo.Session {
  if session == nil {
    var err error

    session, err = mgo.DialWithInfo(&mgo.DialInfo{
      Addrs:    []string{AppConfig.MongoDBHost},
      Username: AppConfig.DBUser,
      Password: AppConfig.DBPwd,
      Timeout:  60 * time.Second,
    })
    if err != nil {
      log.Fatalf("[GetSession]: %s\n", err)
    }
  }
}

func createDbSession() {
  var err error

  session, err = mgo.DialWithInfo(&mgo.DialInfo{
    Username: AppConfig.DBUser,
    Password: AppConfig.DBPwd,
    Timeout:  60 * time.Second,
  })
  if err != nil {
    log.Fatalf("[GetSession]: %s\n", err)
  }
}
