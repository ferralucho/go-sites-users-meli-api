package miapi

type ListingType struct {
	Id 			string `json:"id""`
	NotAvailableCategories		[]string `json:"not_available_in_categories"`
}
