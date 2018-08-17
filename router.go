package bearing

// Router struct holds all defined routes
type Router struct {
	routes map[string]map[string]Route
}

// Route struct holds all information about a defined route
type Route struct {
	Method string
	Path   string
}
