
package arango

import (
  "testing"
)

const (
  C_Mother = iota
  C_Father
  C_Child
)

type Relationship struct{
  Type  int
}

func TestEdge(t *testing.T) {
  var rel Relationship

  ar := _connectDatabase(t)

  db := _createDatabase(ar, t)
  if db == nil {
    return
  }

  collec := db.Collection("vaches")
  err := collec.CreateAuto()
  if err != nil {
    if err, ok := err.(StandardError); ok == true {
      if err.Code != E_Duplicate {
        t.Error("create collection : expected nil, got", err)
        _dropDatabase(db, t)
        return
      }
    } else {
      t.Error("create collection : expected nil, got", err)
      _dropDatabase(db, t)
      return
    }
  }

  collec2 := db.Collection("taureaux")
  err = collec2.CreateAuto()
  if err != nil {
    if err, ok := err.(StandardError); ok == true {
      if err.Code != E_Duplicate {
        t.Error("create collection : expected nil, got", err)
        _dropDatabase(db, t)
        return
      }
    } else {
      t.Error("create collection : expected nil, got", err)
      _dropDatabase(db, t)
      return
    }
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
  err = docf.Create("taureaux", Vache{
    Name: "Coquelicot",
    Age: 4,
    Weight: 390,
  })
  if err != nil {
    t.Error("create document : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }

  edgec1 := db.NewEdge()
  edgec1.From = docc.Handle
  edgec1.To = docm.Handle
  edgec1.Opt.CreateCollection = true
  err = edgec1.Create("relatives", Relationship{C_Child})
  if err != nil {
    t.Error("create edge : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }

  edgec2 := db.NewEdge()
  edgec2.From = docc.Handle
  edgec2.To = docf.Handle
  err = edgec2.Create("relatives", Relationship{C_Child})
  if err != nil {
    t.Error("create edge : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }

  edgep1 := db.NewEdge()
  edgep1.From = docm.Handle
  edgep1.To = docc.Handle
  edgep1.Opt.CreateCollection = true
  err = edgep1.Create("relatives", Relationship{C_Mother})
  if err != nil {
    t.Error("create edge : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }

  edgep2 := db.NewEdge()
  edgep2.From = docf.Handle
  edgep2.To = docc.Handle
  err = edgep2.Create("relatives", Relationship{C_Father})
  if err != nil {
    t.Error("create edge : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }

  edge, err := db.GetEdge(edgep2.Handle)
  if err != nil {
    t.Error("get edge : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }
  if edge.Key != edgep2.Key {
    t.Error("get edge : expected", edgep2.Key, ", got", edge.Key)
    _dropDatabase(db, t)
    return
  }
  if edge.Rev != edgep2.Rev {
    t.Error("get edge : expected", edgep2.Rev, ", got", edge.Rev)
    _dropDatabase(db, t)
    return
  }
  if edge.From != edgep2.From {
    t.Error("get edge : expected", edgep2.From, ", got", edge.From)
    _dropDatabase(db, t)
    return
  }
  if edge.To != edgep2.To {
    t.Error("get edge : expected", edgep2.To, ", got", edge.To)
    _dropDatabase(db, t)
    return
  }

  err = edge.Read(&rel)
  if err != nil {
    t.Error("read edge : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }
  if rel.Type != C_Father {
    t.Error("read edge : expected", C_Father, ", got", rel.Type)
    _dropDatabase(db, t)
    return
  }

  edges, err := docc.Edges("relatives")
  if err != nil {
    t.Error("get edges : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }
  if len(edges) != 4 {
    t.Error("get edges : expected 4, got", len(edges))
    _dropDatabase(db, t)
    return
  }
  if edges[0].From != docm.Handle {
    t.Error("get edges : expected", docm.Handle, ", got", edges[0].From)
    _dropDatabase(db, t)
    return
  }

  edges, err = docc.EdgesIn("relatives")
  if err != nil {
    t.Error("get edges in : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }
  for _, e := range edges {
    if e.To != docc.Handle {
      t.Error("get edges in : expected", docc.Handle, ", got", e.To)
      _dropDatabase(db, t)
      return
    }
  }

  err = edge.Delete()
  if err != nil {
    t.Error("delete edge : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }
  edge, err = db.GetEdge(edge.Handle)
  if err != nil {
    if err, ok := err.(StandardError); ok == true {
      if err.Code != E_NotFound {
        t.Error("get edge : expected nil, got", err)
        _dropDatabase(db, t)
        return
      }
    } else {
      t.Error("get edge : expected nil, got", err)
      _dropDatabase(db, t)
      return
    }
  }

  _dropDatabase(db, t)
}
