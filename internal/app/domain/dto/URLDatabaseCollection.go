package dto

// URLDatabaseCollection - entity to pass a slice with URLs
type URLDatabaseCollection struct {
	collection []DatabaseURLInterface
}

func GetURLDatabaseCollection() URLDatabaseCollection {
	return URLDatabaseCollection{}
}

func (URLDatabaseCollection *URLDatabaseCollection) AppendURL(url DatabaseURLInterface) {
	URLDatabaseCollection.collection = append(URLDatabaseCollection.collection, url)
}

func (URLDatabaseCollection URLDatabaseCollection) GetCollection() []DatabaseURLInterface {
	return URLDatabaseCollection.collection
}
