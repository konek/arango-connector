
package arango

import (
  "fmt"
  "encoding/json"
)

type EdgesRaw struct{
  Edges []Raw `json:"edges"`
}

func (d Document) Edges(collec string) ([]Edge, error) {
  return d.edges(collec, "")
}

func (d Document) EdgesIn(collec string) ([]Edge, error) {
  return d.edges(collec, "in")
}

func (d Document) EdgesOut(collec string) ([]Edge, error) {
  return d.edges(collec, "out")
}

func (d Document) edges(collec string, dir string) ([]Edge, error) {
  var e StandardError
  url := fmt.Sprintf("edges/%s?vertex=%s", collec, d.Handle)
  var raw EdgesRaw

  if len(dir) != 0 {
    url += "&direction=" + dir
  }
  req := d.NewRequest("GET", url)
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
  err = req.Read(&raw)
  if err != nil {
    return nil, err
  }
  ret := make([]Edge, len(raw.Edges))
  for k, r := range raw.Edges {
    ret[k] = *d.NewEdge()
    ret[k].Body = r
    err = json.Unmarshal(r, &ret[k])
    if err != nil {
      return ret, nil
    }
  }
  return ret, nil
}

