package internal

type (
	ID 		string
	Status	uint8
)

const (
	PAID		Status = 1 << iota
	PENDING
	FAILED
	UNKNOWN
)
func validate(invoiceID ID, status Status) bool {

	return !(invoiceID == "" || len(invoiceID) != 8 || status > FAILED)
}