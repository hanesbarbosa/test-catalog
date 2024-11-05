package catalog

// Item is the basic catalog component.
type Item struct {
	id    int     // Unique identification.
	name  string  // Name.
	price float32 // Price.
}

func NewItem() (*Item, error) {
	return nil, nil
}
