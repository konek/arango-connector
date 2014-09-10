
package arango

import (
  "fmt"
)

func (d Database) Drop() error {
  var e StandardError

  url := fmt.Sprintf("database/%s", d.Arango.Db)
  req := d.NewRequest("DELETE", url)
  req.System = true
  resp, err := req.Do()
  if err != nil {
    return err
  }
  err = req.Read(&e)
  if err != nil {
    return e
  }
  if e.Err == true {
    return e
  }
  if resp.StatusCode != 200 {
    return UnknownError{resp.StatusCode}
  }
  return nil
}

