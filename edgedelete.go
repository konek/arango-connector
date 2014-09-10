
package arango

import (
  "fmt"
)

func (e Edge) Delete() error {
  var er StandardError
  url := fmt.Sprintf("edge/%s", e.Handle)

  if e.Opt.WaitForSync == true {
    url += "?waitForSync=true"
  }
  req := e.NewRequest("DELETE", url)
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
  if resp.StatusCode != 200 && resp.StatusCode != 202 {
    return UnknownError{resp.StatusCode}
  }
  return nil
}

