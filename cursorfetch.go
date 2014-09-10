
package arango

func (c *Cursor) Fetch() error {
  var e StandardError
  url := "cursor/" + c.Id

  req := c.NewRequest("PUT", url)
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
  if resp.StatusCode != 200 {
    return UnknownError{resp.StatusCode}
  }
  err = req.Read(c)
  if err != nil {
    return err
  }
  if len(c.Result_) > 0 {
    c.HasMore = true
  }
  c.pos = 0
  return nil
}

