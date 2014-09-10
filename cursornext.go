
package arango

func (c *Cursor) Next() (*Document, error) {
  var ret *Document

  if c.pos >= len(c.Result_) {
    if c.HasMore_ == false {
      // Error no more results
      return nil, StandardError{
        Err: true,
        Code: E_NotFound,
        ErrorNum: -1,
        ErrorMessage: "No more results",
      }
    }
    // Next Batch
    err := c.Fetch()
    if err != nil {
      return nil, err
    }
    return c.Next()
  }
  ret = &c.Result_[c.pos].Document
  c.pos += 1
  if c.pos >= len(c.Result_) && c.HasMore_ == false {
    c.HasMore = false
  }
  return ret, nil
}

