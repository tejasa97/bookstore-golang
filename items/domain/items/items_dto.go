package items

type Item struct {
	ID                string      `json:"id"`
	Seller            int64       `json:"seller"`
	Title             string      `json:"title"`
	Description       Description `json:"description"`
	Pictures          []Picture   `json:"pictures"`
	Video             string      `json:"video"`
	Price             float32     `json:"price"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Status            string      `json:"status"`
}

type Description struct {
	PlainText string `json:"plain_test"`
	HTML      string `json:"html"`
}

type Picture struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}
