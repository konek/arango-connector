
package arango

type ParsedQuery struct{
  Vars        []string  `json:"bindVars"`
  Collections []string  `json:"collections"`
}

func (q *Query) Parse() (ParsedQuery, error) {
  var e StandardError
  var ret ParsedQuery

  req := q.NewRequest("POST", "query")
  req.Data = *q
  resp, err := req.Do()
  if err != nil {
    return ret, err
  }
  err = req.Read(&e)
  if err != nil {
    return ret, err
  }
  if e.Err == true {
    return ret, e
  }
  if resp.StatusCode != 200 {
    return ret, UnknownError{resp.StatusCode}
  }
  err = req.Read(&ret)
  if err != nil {
    return ret, err
  }
  return ret, nil
}

