
package arango

type CreateCollectionQuery struct{
  Name        string    `json:"name" binding:"required"`
  WaitForSync bool      `json:"waitForSync"`
  DoCompact   bool      `json:"doCompact"`
  JournalSize int       `json:"journalSize,omitempty"`
  IsSystem    bool      `json:"isSystem"`
  IsVolatile  bool      `json:"isVolatile"`
  Type        int       `json:"type,omitempty"`
  Shards      int       `json:"numberOfShards,omitempty"`
  ShardKeys   []string  `json:"shardKeys,omitempty"`

  KeyOptions  []KeyOption `json:"keyOptions,omitempty"`
}

const (
  CollectionTypeDocument = 2
  CollectionTypeEdges = 3
)

func (c *Collection) CreateAuto() error {
  c.WaitForSync = false
  c.DoCompact = true
  c.IsSystem = false
  c.IsVolatile = false
  if c.Type == 0 {
    c.Type = 2
  }
  if c.Shards == 0 {
    c.Shards = 1
    c.ShardKeys = nil
  }
  return c.Create()
}

func (c *Collection) Create() error {
  var e StandardError

  req := c.NewRequest("POST", "collection")
  req.Data = c
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
  err = req.Read(&c)
  if err != nil {
    return err
  }
  return nil
}

