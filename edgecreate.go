
package arango

import (
  "fmt"
)

func (e *Edge) Create(collec string, obj interface{}) error {
  var er StandardError
  url := fmt.Sprintf("edge?collection=%s", collec)

  if e.Opt.CreateCollection == true {
    url += "&createCollection=true"
  }
  if e.Opt.WaitForSync == true {
    url += "&waitForSync=true"
  }
  url += "&from=" + e.From
  url += "&to=" + e.To

  req := e.NewRequest("POST", url)
  req.Data = obj
  resp, err := req.Do()
  if err != nil {
    return err
  }
  err = req.Read(&er)
  if err != nil {
    return err
  }
  if er.Err == true {
    return er
  }
  if resp.StatusCode != 201 && resp.StatusCode != 202 {
    return UnknownError{resp.StatusCode}
  }
  err = req.Read(e)
  if err != nil {
    return err
  }
  return nil
}

func (e *Edge) CreateEmpty(collec string) error {
  return e.Create(collec, Empty{})
}

