package theservice

var allEntities = []TheService{
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
	{Title: "five"},
	{Title: "six"},
	{Title: "seven"},
	{Title: "eight"},
	{Title: "nine"},
	{Title: "ten"},
	{Title: "eleven"},
	{Title: "twelve"},
	{Title: "thirteen"},
	{Title: "fourteen"},
}

type TheService struct {
	Title string
}

func (service *TheService) String() string {
	return service.Title
}
