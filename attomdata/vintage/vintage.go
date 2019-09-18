package vintage

//Vintage ...
type Vintage struct {
	LastModified string `json:"lastModified"`
	PubDate      string `json:"pubDate"`
}

//NewVintage ...
func NewVintage(lastModified string, pubDate string) Vintage {
	return Vintage{
		LastModified: lastModified,
		PubDate:      pubDate,
	}
}
