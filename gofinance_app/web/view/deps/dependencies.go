package deps

const HTMX_WS = "htmx_ws"

var dependencyMap map[string]string

func initializeMap() {
	dependencyMap = make(map[string]string)
	dependencyMap[HTMX_WS] = "https://unpkg.com/htmx.org/dist/ext/ws.js"
}

func GetLink(name string) string {
	if dependencyMap == nil {
		initializeMap()
	}
	return dependencyMap[name]
}
