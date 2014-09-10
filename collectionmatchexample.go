
package arango

import (
  "encoding/json"
)

type MatchExampleQuery struct{
  Collection  string      `json:"collection"`
  Example     interface{} `json:"example"`
}

type Raw []byte

type MatchExampleResp struct{
  Document  Raw  `json:"document"`
}

func (r *Raw) UnmarshalJSON(raw []byte) error {
  *r = raw
  return nil
}

func (c Collection) MatchExample(example interface{}) (*Document, error) {
  var e StandardError
  var r MatchExampleResp
  doc := c.NewDocument()

  q := MatchExampleQuery{
    Collection: c.Name,
    Example: example,
  }

  req := c.NewRequest("PUT", "simple/first-example")
  req.Data = q
  resp, err := req.Do()
  if err != nil {
    return nil, err
  }
  err = req.Read(&e)
  if err != nil {
    return nil, err
  }
  if e.Err == true {
    return nil, e
  }
  if resp.StatusCode != 200 {
    return nil, UnknownError{resp.StatusCode}
  }
  err = json.Unmarshal(req.Body, &r)
  if err != nil {
    return nil, err
  }
  err = json.Unmarshal(r.Document, &doc)
  if err != nil {
    return nil, err
  }
  doc.Body = r.Document
  return doc, nil
}

