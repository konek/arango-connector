
package arango

import (
  "testing"
)

func _connectDatabase(t *testing.T) *Arango {
  ar := New("http://localhost:8529")
  ar.User = "root"
  ar.Passwd = "kidoo blur carcrash airplane"
  return ar
}

func _createDatabase(ar *Arango, t *testing.T) *Database {
  db := ar.NewDatabase("konek_test")
  err := db.CreateAuto()
  if err != nil {
    t.Error("create db : expected nil, got ", err)
    return nil
  }
  return db
}

func _dropDatabase(db *Database, t *testing.T) {
  err := db.Drop()
  if err != nil {
    t.Error("drop db : expected nil, got ", err)
  }
}

