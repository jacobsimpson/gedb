package metadata

var cachedMetadata = map[string]*storage.TableMetadata{
	"objects": storage.TableMetadata{
		TableId:   0,
		TableName: "objects",
		FirstPage: 0,
	},
}

func TableMetadata(tableName string) *storage.TableMetadata {
	return cachedMetadata[tableName]
}
