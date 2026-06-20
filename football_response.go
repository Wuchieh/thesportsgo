package thesportsgo

type FootballCategoryResponseData struct {
	// Category id
	ID string `json:"id"`
	// Category Name
	Name string `json:"name"`
	// Update time
	UpdatedAt int `json:"updated_at"`
}

type FootballCountryResponseData struct {
	// Category id
	CategoryID string `json:"category_id"`
	// Country/Region id
	ID string `json:"id"`
	// Country/Region logo
	Logo string `json:"logo"`
	// Country/Region Name
	Name string `json:"name"`
	// Update time
	UpdatedAt int64 `json:"updated_at"`
}
