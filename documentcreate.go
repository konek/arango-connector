
package arango

import (
  "fmt"
)

func (d *Document) Create(collec string, doc interface{}) error {
  var e StandardError
  url := fmt.Sprintf("document?collection=%s", collec)

  if d.Opt.WaitForSync == true {
    url += "&waitForSync=true"
  }
  if d.Opt.CreateCollection == true {
    url += "&createCollection=true"
  }
  req := d.NewRequest("POST", url)
  req.Data = doc
  resp, err := req.Do()
  if err != nil {
    return err
  }
  err = req.Read(&e)
  if err != nil {
    return err
  }
  if e.Err == true {
    return e
  }
  if resp.StatusCode != 201 && resp.StatusCode != 202 {
    return UnknownError{resp.StatusCode}
  }
  err = req.Read(d)
  if err != nil {
    return err
  }
  return nil
}

