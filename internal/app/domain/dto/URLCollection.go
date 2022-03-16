package dto

import (
	"github.com/PanovAlexey/url_carver/internal/app/domain"
)

// URLCollection - entity to pass a slice with URLs
type URLCollection struct {
	collection []domain.URLInterface
}

func GetURLCollection() *URLCollection {
	return &URLCollection{}
}

func (URLCollection *URLCollection) AppendURL(url domain.URLInterface) {
	URLCollection.collection = append(URLCollection.collection, url)
}

func (URLCollection *URLCollection) GetCollection() []domain.URLInterface {
	return URLCollection.collection
}
