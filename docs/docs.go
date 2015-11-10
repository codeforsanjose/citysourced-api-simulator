package docs

type Page struct {
	Title string
	Body  string
}

var (
	Home = &Page{
		Title: "CitySourced Test API",
	    Body: `
			<p>This is the documentation for the City Sourced test API System.
			Enter URL &#8220;docs&#8221; for more information.</p>
		`,
	}
	
	Detail = &Page{
		Title: "Instructions",
		Body: `
			<p>This is detailed documentation for the City Sourced test API System.</p>
		`,
	}
)
