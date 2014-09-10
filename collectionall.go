
package arango

type AllQuery struct{
  Collection  string  `json:"collection"`
  BatchSize   int     `json:"batchSize,omitempty"`
}

func (c Collection) All() (*Cursor, error) {
  return c.AllBatch(4)
}

func (c Collection) AllBatch(batch int) (*Cursor, error) {
  var e StandardError
  ret := c.Cursor()
  q := AllQuery{
    Collection: c.Name,
    BatchSize: batch,
  }

  req := c.NewRequest("PUT", "simple/all")
  req.Data = q
  resp, err := req.Do()
  if err != nil {
    return nil, err
  }
  err = req.Read(&e)
  if err != nil {
    return nil, err
  }
  if e.Err == true {
    return nil, e
  }
  if resp.StatusCode != 201 {
    return nil, UnknownError{resp.StatusCode}
  }
  err = req.Read(ret)
  if err != nil {
    return nil, err
  }
  if len(ret.Result_) > 0 {
    ret.HasMore = true
  }
  return ret, nil
}

