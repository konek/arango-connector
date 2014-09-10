
package arango

type Database struct{
  Arango
}

func (a Arango) NewDatabase(dbname string) *Database {
  ret := &Database{
    a,
  }
  ret.Arango.Db = dbname

  return ret
}

