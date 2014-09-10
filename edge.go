
package arango

import (
  "encoding/json"
)

type Edge struct{
  *Database         `json:"-"`
  Handle    string  `json:"_id"`
  Rev       string  `json:"_rev"`
  Key       string  `json:"_key"`
  From      string  `json:"_from"`
  To        string  `json:"_to"`

  Opt       Options `json:"-"`
  Body      []byte  `json:"-"`
}

func (a Arango) NewEdge() *Edge {
  return &Edge{
    Database: a.Use(a.Db),
    Opt: OptionsDefaults(),
  }
}

func (e Edge) Read(obj interface{}) error {
  return json.Unmarshal(e.Body, obj)
}

