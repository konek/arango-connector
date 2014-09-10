
package arango

import (
  "fmt"
)

func (d *Document) Fetch() error {
  var e StandardError
  url := fmt.Sprintf("document/%s", d.Handle)

  req := d.NewRequest("GET", url)
  resp, err := req.Do()
  if err != nil {
    return err
  }
  err = req.Read(&e)
  if err == nil {
    if e.Err == true {
      return e
    }
  }
  if resp.StatusCode != 200 {
    return UnknownError{resp.StatusCode}
  }
  d.Body = req.Body
  return nil
}

