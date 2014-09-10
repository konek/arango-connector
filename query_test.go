
package arango

import (
  "testing"
)

func TestQueryExplain(t *testing.T) {
  ar := _connectDatabase(t)

  db := _createDatabase(ar, t)
  if db == nil {
    return
  }

  collec := db.Collection("products")
  err := collec.CreateAuto()
  if err != nil {
    t.Error("create collection : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }

  q := db.Query("FOR p IN products FILTER p.id == @id LIMIT 2 RETURN p.name")
  q.Bind("id", 3)
  expl, err := q.Explain()
  if err != nil {
    t.Error("explain query : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }
  if len(expl) != 3 {
    t.Error("explain query : expected 3, got", len(expl))
    _dropDatabase(db, t)
    return
  }
  if expl[0].Id != 1 {
    t.Error("explain query : expected 1, got", expl[0].Id)
    _dropDatabase(db, t)
    return
  }
  if expl[0].Type != "for" {
    t.Error("explain query : expected 'for', got", expl[0].Type)
    _dropDatabase(db, t)
    return
  }
  if expl[0].LoopLevel != 1 {
    t.Error("explain query : expected 1, got", expl[0].LoopLevel)
    _dropDatabase(db, t)
    return
  }
  if expl[0].ResultVariable != "p" {
    t.Error("explain query : expected 'p', got", expl[0].ResultVariable)
    _dropDatabase(db, t)
    return
  }
  if expl[0].Limit != true {
    t.Error("explain query : expected true, got", expl[0].Limit)
    _dropDatabase(db, t)
    return
  }
  if expl[0].Expr.Type != "collection" {
    t.Error("explain query : expected 'collection', got", expl[0].Expr.Type)
    _dropDatabase(db, t)
    return
  }
  if expl[0].Expr.Val != "products" {
    t.Error("explain query : expected 'products', got", expl[0].Expr.Val)
    _dropDatabase(db, t)
    return
  }

  _dropDatabase(db, t)
}

func TestQueryParse(t *testing.T) {
  ar := _connectDatabase(t)

  db := _createDatabase(ar, t)
  if db == nil {
    return
  }

  collec := db.Collection("products")
  err := collec.CreateAuto()
  if err != nil {
    t.Error("create collection : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }

  q := db.Query("FOR p IN products FILTER p.id == @id LIMIT 2 RETURN p.name")
  q.Bind("id", 3)
  parsed, err := q.Parse()
  if err != nil {
    t.Error("parse query : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }
  if len(parsed.Vars) != 1 {
    t.Error("parse query : expected 1, got", len(parsed.Vars))
    _dropDatabase(db, t)
    return
  }
  if parsed.Vars[0] != "id" {
    t.Error("parse query : expected 'id', got", parsed.Vars[0])
    _dropDatabase(db, t)
    return
  }
  if len(parsed.Collections) != 1 {
    t.Error("parse query : expected 1, got", len(parsed.Collections))
    _dropDatabase(db, t)
    return
  }
  if parsed.Collections[0] != "products" {
    t.Error("parse query : expected 'products', got", parsed.Collections[0])
    _dropDatabase(db, t)
    return
  }

  _dropDatabase(db, t)
}

func TestQuery(t *testing.T) {
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

  q := db.Query("FOR v IN vaches RETURN v")
  cur, err := q.Run()
  if err != nil {
    t.Error("query run : expected nil, got", err)
    _dropDatabase(db, t)
    return
  }
  if cur.Count != 3 {
    t.Error("query run : expected 3, got", cur.Count)
    _dropDatabase(db, t)
    return
  }
  for i := 0; cur.HasMore == true; i++ {
    var v Vache

    doc, err := cur.Next()
    if err != nil {
      t.Error("cursor next : expected nil, got", err)
      _dropDatabase(db, t)
      return
    }
    if doc == nil {
      t.Error("cursor next : expected document, got nil")
      _dropDatabase(db, t)
      return
    }
    err = doc.Read(&v)
    if err != nil {
      t.Error("document read : expected nil, got", err)
      continue
    }
    if cur.HasMore == false && i != 2 {
      t.Error("cursor next : expected 2, got", i)
      _dropDatabase(db, t)
      return
    }
  }
  _dropDatabase(db, t)
}

