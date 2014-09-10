
package arango

type Query struct{
  Database          `json:"-"`
  Query     string  `json:"query"`
  Vars  map[string]interface{}  `json:"bindVars,omitempty"`
  Count     bool    `json:"count"`
  Batch     int     `json:"batchSize"`
}

func (db Database) Query(q string) Query {
  return Query{
    Database: db,
    Query: q,
    Count: true,
    Batch: 2,
    Vars: make(map[string]interface{}),
  }
}

func (q *Query) Bind(k string, v interface{}) {
  q.Vars[k] = v
}

