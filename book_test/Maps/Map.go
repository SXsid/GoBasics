package maps

type Dict map[string]string

func (d Dict) Search(key string) string {
	return d[key]
}
