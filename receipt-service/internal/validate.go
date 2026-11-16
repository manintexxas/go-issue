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

	if invoiceID == "" || status > UNKNOWN {
		return false
	}

	return true
}