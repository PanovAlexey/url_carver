package dto

import "github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"

// URLCollection - entity to pass a slice with URLs
type URLCollection struct {
	collection []url.URL
}

func GetURLCollection() *URLCollection {
	return &URLCollection{}
}

func (URLCollection *URLCollection) AppendURL(url url.URL) {
	URLCollection.collection = append(URLCollection.collection, url)
}

func (URLCollection *URLCollection) GetCollection() []url.URL {
	return URLCollection.collection
}
