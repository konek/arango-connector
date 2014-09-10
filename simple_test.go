
package arango

import (
  "testing"
)

type VacheExample struct{
  Age int
}

func TestSimpleReturnAll(t *testing.T) {
  ar := _connectDatabase(t)

  db := _createDatabase(ar, t)
  if db == nil {
    return
  }

  collec := db.Collection("vaches")
  err := collec.CreateAuto()
  if err != nil {
    t.Error("create collection : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }

  docc := db.NewDocument()
  err = docc.Create("vaches", Vache{
    Name: "Paquerette",
    Age: 1,
    Weight: 180,
  })
  if err != nil {
    t.Error("create document : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }

  docm := db.NewDocument()
  err = docm.Create("vaches", Vache{
    Name: "Tulipe",
    Age: 3,
    Weight: 250,
  })
  if err != nil {
    t.Error("create document : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }

  docf := db.NewDocument()
  err = docf.Create("vaches", Vache{
    Name: "Coquelicot",
    Age: 4,
    Weight: 390,
  })
  if err != nil {
    t.Error("create document : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }

  cur, err := collec.AllBatch(2)
  if err != nil {
    t.Error("collection all : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }
  if cur.HasMore != true {
    t.Error("collection all : expected true, got", cur.HasMore)
    _dropDatabase(db, t)
    return
  }
  if cur.HasMore_ != true {
    t.Error("collection all : expected true, got", cur.HasMore_)
    _dropDatabase(db, t)
    return
  }
  if cur.Count != 3 {
    t.Error("collection all : expected 3, got", cur.Count)
    _dropDatabase(db, t)
    return
  }
  if len(cur.Result_) != 2 {
    t.Error("collection all : expected 2, got", len(cur.Result_))
    _dropDatabase(db, t)
    return
  }

  _dropDatabase(db, t)
}

func TestSimpleMatch(t *testing.T) {
  ar := _connectDatabase(t)

  db := _createDatabase(ar, t)
  if db == nil {
    return
  }

  collec := db.Collection("vaches")
  err := collec.CreateAuto()
  if err != nil {
    t.Error("create collection : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }

  docc := db.NewDocument()
  err = docc.Create("vaches", Vache{
    Name: "Paquerette",
    Age: 1,
    Weight: 180,
  })
  if err != nil {
    t.Error("create document : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }

  docm := db.NewDocument()
  err = docm.Create("vaches", Vache{
    Name: "Tulipe",
    Age: 1,
    Weight: 250,
  })
  if err != nil {
    t.Error("create document : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }

  docf := db.NewDocument()
  err = docf.Create("vaches", Vache{
    Name: "Coquelicot",
    Age: 4,
    Weight: 390,
  })
  if err != nil {
    t.Error("create document : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }

  eg := VacheExample{
    Age: 1,
  }
  cur, err := collec.Match(eg)
  if err != nil {
    t.Error("collection match : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }
  if cur.HasMore != true {
    t.Error("collection all : expected true, got", cur.HasMore)
    _dropDatabase(db, t)
    return
  }
  if cur.HasMore_ != false {
    t.Error("collection all : expected false, got", cur.HasMore_)
    _dropDatabase(db, t)
    return
  }
  if cur.Count != 2 {
    t.Error("collection all : expected 2, got", cur.Count)
    _dropDatabase(db, t)
    return
  }
  if len(cur.Result_) != 2 {
    t.Error("collection all : expected 2, got", len(cur.Result_))
    _dropDatabase(db, t)
    return
  }

  _dropDatabase(db, t)
}

