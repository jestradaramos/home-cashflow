package mortgage

type Mortgage struct {
	MaxPrice    int
	Downpayment float32
	Monthly     float32
	PMI         float32
	HOI         float32
}

// All functions needed for the Mortgage object
func NewMortgage(price int, downpayment, monthly, pmi, hoi float32) *Mortgage {
	return &Mortgage{price, downpayment, monthly, pmi, hoi}
}

// These setters need to be able to modify off each other if possible
func (m *Mortgage) SetPrice(price int) {
	m.MaxPrice = price
}
func (m *Mortgage) SetDownPayment(dp float32) {
	m.Downpayment = dp
}
func (m *Mortgage) SetMonthly(monthly float32) {
	m.Monthly = monthly
}
func (m *Mortgage) SetPMI(pmi float32) {
	m.PMI = pmi
}
func (m *Mortgage) SetHOI(hoi float32) {
	m.HOI = hoi
}
