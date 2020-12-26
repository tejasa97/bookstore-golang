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

// func (i *Item) Validate() *rest_errors.RestErr {

// 	if i.Seller == 0 {
// 		return rest_errors.NewBadRequestError("invalid seller id")
// 	}
// 	if i.Title == "" {
// 		return rest_errors.NewBadRequestError("invalid title")
// 	}
// 	if i.Description == "" {
// 		return rest_errors.NewBadRequestError("invalid description")
// 	}
// 	if i.Price == 0 {
// 		return rest_errors.NewBadRequestError("invalid price")
// 	}
// 	if i.AvailableQuantity == 0 {
// 		return rest_errors.NewBadRequestError("invalid available quantity")
// 	}
// 	if i.SoldQuantity == 0 {
// 		return rest_errors.NewBadRequestError("invalid seller id")
// 	}
// 	if i.Status == "" {
// 		return rest_errors.NewBadRequestError("invalid status")
// 	}

// 	return nil
// }
