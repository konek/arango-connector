
package arango

import (
  "fmt"
)

func (a Arango) GetDocument(handle string) (*Document, error) {
  var e StandardError
  doc := a.NewDocument()
  url := fmt.Sprintf("document/%s", handle)

  req := a.NewRequest("GET", url)
  resp, err := req.Do()
  if err != nil {
    return nil, err
  }
  err = req.Read(&e)
  if err == nil {
    if e.Err == true {
      return nil, e
    }
  }
  if resp.StatusCode != 200 {
    return nil, UnknownError{resp.StatusCode}
  }
  doc.Body = req.Body
  err = req.Read(&doc)
  return doc, err
}

func (c Collection) GetDocument(Key string) (*Document, error) {
  handle := fmt.Sprintf("%s/%s", c.Name, Key)
  return c.Database.GetDocument(handle)
}

