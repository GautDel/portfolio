package handlers

type Project struct {
	Name 		string
	CodedDate   int
	ProjectType string 
	Url         string 
}

type Projects []Project

type PostSummary struct {
	Title   string
	Created string
	Slug    string
}

type PostSummaries []PostSummary

type Image struct {
	Image 	 string
	Alt 	 string
}

type Images []Image

type RichText struct {
	RichText string
}

type Page struct {
	Page        string
	RichText    string
	Description string
}
