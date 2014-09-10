
package arango

import (
  "testing"
)

func TestDb(t *testing.T) {
  ar := New("http://localhost:8529")
  ar.User = "root"
  ar.Passwd = "kidoo blur carcrash airplane"

  db := ar.NewDatabase("konek_test")
  err := db.CreateAuto()
  if err != nil {
    if err, ok := err.(StandardError); ok == true {
      if err.Code != E_Duplicate {
        t.Error("create db : expected nil, got", err)
        _dropDatabase(db, t)
        return
      }
    } else {
      t.Error("create db : expected nil, got", err)
      _dropDatabase(db, t)
    }
  }

  err = db.Drop()
  if err != nil {
    t.Error("drop db : expected nil, got", err)
  }
}

