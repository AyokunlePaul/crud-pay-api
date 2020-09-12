package product

type Response struct {
	Status bool
	Message string
	Data interface{}
}

type Meta struct {
	Total, Skipped, PerPage, Page, PageCount int64
}