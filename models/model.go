package models

type News struct {
	Headline string
	Content  string
}

var SomeNews = []News{
	{
		Headline: "news1 headline",
		Content:  "hello",
	},
	{
		Headline: "news2 headline",
		Content:  "ur not",
	},
	{
		Headline: "news3 headline",
		Content:  "nice",
	},
}
