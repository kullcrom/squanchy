package vintage

type Vintage struct {
	LastModified string
	PubDate      string
}

func NewVintage(lastModified string, pubDate string) Vintage {
	return Vintage{
		LastModified: lastModified,
		PubDate:      pubDate,
	}
}
