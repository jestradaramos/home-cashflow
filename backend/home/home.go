package home

type Home struct {
	Zip     int
	Address string
	Beds    int
	Baths   float32
	Price   int
	Taxes   float32
	Rent    int
	Utils   float32
}

// Place of interest (e.g. work, frequent spots, parents, work)
type POI struct {
	Address string
	Method  string
}

// All the functions for the Home Objects
func NewHome(address string, beds int, baths float32, price int, rent int, utils float32) *Home {
	zip := getZip(address)
	taxes := getTaxRate(zip)
	return &Home{
		Address: address,
		Zip:     zip,
		Beds:    beds,
		Baths:   baths,
		Price:   price,
		Taxes:   taxes,
		Rent:    rent,
		Utils:   utils,
	}
}
func NewHomeZillow(address string) *Home {
	return getHome(address)
}

// HELPER FUNCTIONS

// Use API or library to get tax rates by zip code
// Zillow seems promosing providing hitsorical tax data
func getTaxRate(zip int) float32 {
	return 0
}

func getHome(address string) *Home {
	// run zillow api
	return &Home{}
}

// Regex string parser
func getZip(address string) int {
	return 0
}
