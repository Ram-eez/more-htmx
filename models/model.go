package models

type News struct {
	Headline string
	Content  string
}

var SomeNews = []News{
	{
		Headline: "news1 headline",
		Content:  "something something something",
	},
	{
		Headline: "news2 headline",
		Content:  "something something something",
	},
	{
		Headline: "news3 headline",
		Content:  "something something something",
	},
}
