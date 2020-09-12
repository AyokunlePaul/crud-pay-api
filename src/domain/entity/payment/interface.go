package payment

type reader interface {
	GetWithId(string, string) (*Transaction, *Response)
	Get(string) ([]Transaction, *Response)
}

type writer interface {
	Create()
	Update()
}
