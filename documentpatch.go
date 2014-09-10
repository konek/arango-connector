
package arango

import (
  "fmt"
)

func (d Document) Patch(obj interface{}) error {
  var e StandardError

  url := fmt.Sprintf("document/%s?", d.Handle)
  if d.Opt.KeepNull == true {
    url += "keepNull=true"
  } else {
    url += "keepNull=false"
  }
  if d.Opt.WaitForSync == true {
    url += "&waitForSync=true"
  } else {
    url += "&waitForSync=false"
  }
  req := d.NewRequest("PATCH", url)
  req.Data = obj
  resp, err := req.Do()
  if err != nil {
    return err
  }
  err = req.Read(&e)
  if err != nil {
    return err
  }
  if e.Err == true {
    return  e
  }
  if resp.StatusCode != 201 && resp.StatusCode != 202 {
    return UnknownError{resp.StatusCode}
  }
  return nil
}

