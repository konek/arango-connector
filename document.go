
package arango

import (
  "encoding/json"
)

type Document struct{
  *Database         `json:"-"`

  Handle    string  `json:"_id"`
  Rev       string  `json:"_rev"`
  Key       string  `json:"_key"`

  Opt       Options `json:"-"`
  Body      []byte  `json:"-"`
}

type JSONDocument struct{
  Document
}

func (a Arango) NewDocument() *Document {
  return &Document{
    Database: a.Use(a.Db),
    Opt: OptionsDefaults(),
  }
}

func (d Document) Read(obj interface{}) error {
  return json.Unmarshal(d.Body, obj)
}

func (d *JSONDocument) UnmarshalJSON(raw []byte) error {
  err := json.Unmarshal(raw, &d.Document)
  d.Document.Body = raw
  return err
}

