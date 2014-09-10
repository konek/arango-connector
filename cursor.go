
package arango

type Cursor struct{
  Database                `json:"-"`

  Id        string          `json:"id"`
  HasMore_  bool            `json:"hasMore"`
  Result_   []JSONDocument  `json:"result"`
  pos       int             `json:"-"`

  Count     int             `json:"count"`
  HasMore   bool            `json:"-"`
}

func (d Database) Cursor() *Cursor {
  return &Cursor{
    Database: d,
  }
}

