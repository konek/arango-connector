
package arango

import (
  "net/http"
)

type Arango struct{
  http.Client

  User    string
  Passwd  string
  Addr    string
  Db      string
}

func New(addr string) *Arango {
  return &Arango{
    Addr: addr,
  }
}

func (a Arango) Use(db string) *Database {
  return a.NewDatabase(db)
}

