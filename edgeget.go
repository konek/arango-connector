
package arango

import (
  "fmt"
)

func (a Arango) GetEdge(handle string) (*Edge, error) {
  var e StandardError
  edge := a.NewEdge()
  url := fmt.Sprintf("edge/%s", handle)

  req := a.NewRequest("GET", url)
  resp, err := req.Do()
  if err != nil {
    return nil, err
  }
  err = req.Read(&e)
  if err == nil {
    if e.Err == true {
      return nil, e
    }
  }
  if resp.StatusCode != 200 {
    return nil, UnknownError{resp.StatusCode}
  }
  edge.Body = req.Body
  err = req.Read(&edge)
  return edge, err
}

