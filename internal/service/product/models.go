package product

type Product struct {
	Title string `json:"title"`
}

var allProducts []Product = []Product{
	{Title: "first"},
	{Title: "second"},
	{Title: "3th"},
	{Title: "4th"},
}
