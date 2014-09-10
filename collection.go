
package arango

type KeyOption struct{
  Type          string  `json:"type"`
  AllowUserKeys bool    `json:"allowUserKeys"`
  Increment     string  `json:"increment"`
  Offset        string  `json:"offset"`
}

type Collection struct{
  Database  `json:"-"`

  Id          string      `json:"id"`
  Name        string      `json:"name"`
  WaitForSync bool        `json:"waitForSync"`
  DoCompact   bool        `json:"doCompact"`
  JournalSize int         `json:"journalSize,omitempty"`
  IsSystem    bool        `json:"isSystem"`
  IsVolatile  bool        `json:"isVolatile"`
  Type        int         `json:"type"`
  Shards      int         `json:"numberOfShards,omitempty"`
  ShardKeys   []string    `json:"shardKeys,omitempty"`
  KeyOptions  []KeyOption `json:"keyOptions,omitempty"`

  Status      int         `json:"status"`
}

func (db Database) Collection(name string) Collection {
  return Collection{
    Database: db,
    Name: name,
  }
}

