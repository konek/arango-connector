
package arango

import (
  "fmt"
)

func (d Document) Delete() error {
  var e StandardError
  url := fmt.Sprintf("document/%s", d.Handle)

  req := d.NewRequest("DELETE", url)
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
  if resp.StatusCode != 202 {
    return UnknownError{resp.StatusCode}
  }
  return nil
}

