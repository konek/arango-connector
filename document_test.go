
package arango

import (
  "testing"
)

type Vache struct{
  Name    string  `json:",omitempty"`
  Age     int     `json:",omitempty"`
  Weight  int     `json:",omitempty"`
}

func TestDocument(t *testing.T) {
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
  if collec.Id == "" {
    t.Error("create collection : expected Id, got", collec.Id)
    _dropDatabase(db, t)
    return
  }

  vache := Vache{
    Name: "Paquerette",
    Age: 2,
    Weight: 220,
  }
  doc := db.NewDocument()
  err = doc.Create("vaches", vache)
  if err != nil {
    t.Error("create document : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }
  if doc.Handle == "" {
    t.Error("create document : expected handle, got ''")
    _dropDatabase(db, t)
    return
  }

  if doc.Rev == "" {
    t.Error("create document : expected rev, got ''")
    _dropDatabase(db, t)
    return
  }

  if doc.Key == "" {
    t.Error("create document : expected key, got ''")
    _dropDatabase(db, t)
    return
  }

  doc2, err := db.GetDocument(doc.Handle)
  if err != nil {
    t.Error("get document : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }
  if doc2.Handle != doc.Handle {
    t.Error("get document : expected handle", doc.Handle, ", got", doc2.Handle)
    _dropDatabase(db, t)
    return
  }
  if doc2.Rev != doc.Rev {
    t.Error("get document : expected rev", doc.Rev, ", got", doc2.Rev)
    _dropDatabase(db, t)
    return
  }
  if doc2.Key != doc.Key {
    t.Error("get document : expected key", doc.Key, ", got", doc2.Key)
    _dropDatabase(db, t)
    return
  }

  vache2 := Vache{}
  err = doc2.Read(&vache2)
  if err != nil {
    t.Error("read document : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }
  if vache2.Name != vache.Name {
    t.Error("get document : expected name", vache.Name, ", got", vache2.Name)
    _dropDatabase(db, t)
    return
  }
  if vache2.Age != vache.Age {
    t.Error("get document : expected age", vache.Age, ", got", vache2.Age)
    _dropDatabase(db, t)
    return
  }
  if vache2.Weight != vache.Weight {
    t.Error("get document : expected weight", vache.Weight, ", got", vache2.Weight)
    _dropDatabase(db, t)
    return
  }

  vache2 = Vache{
    Age: 42,
  }
  err = doc2.Patch(vache2)
  if err != nil {
    t.Error("patch document : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }
  err = doc2.Fetch()
  if err != nil {
    t.Error("fetch document : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }
  err = doc2.Read(&vache2)
  if err != nil {
    t.Error("read document : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }
  if vache2.Name != vache.Name {
    t.Error("patch/fetch document : expected name", vache.Name, ", got", vache2.Name)
    _dropDatabase(db, t)
    return
  }
  if vache2.Age != 42 {
    t.Error("patch/fetch document : expected age", vache.Age, ", got", vache2.Age)
    _dropDatabase(db, t)
    return
  }
  if vache2.Weight != vache.Weight {
    t.Error("patch/fetch document : expected weight", vache.Weight, ", got", vache2.Weight)
    _dropDatabase(db, t)
    return
  }

  _dropDatabase(db, t)
}

