package rapid

import (
	"strings"
)

func getParams(path string, rPath string) map[string]string {
	paramLocations := getParamLocations(path)
	requestPath := strings.Split(rPath, "/")
	params := map[string]string{}

	if len(paramLocations) > 0 {
		for i := 0; i < len(requestPath); i++ {
			if val, ok := paramLocations[i]; ok {
				params[val] = requestPath[i]
			}
		}
	}

	return params
}

func getParamLocations(path string) map[int]string {

	routePath := strings.Split(path, "/")

	params := map[int]string{}
	for i := (len(routePath) - 1); i >= 0; i-- {
		dir := routePath[i]
		if dir != "" && dir[0] == ':' {
			params[i] = dir[1:len(routePath[i])]
			routePath = append(routePath[:i], routePath[i+1:]...)
		}
	}
	newPath := strings.Join(routePath, "/")
	if len(params) > 0 {
		newPath += "/"
	}
	return params
}
