package theservice

var allEntities = []TheService{
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
	{Title: "five"},
}

type TheService struct {
	Title string
}