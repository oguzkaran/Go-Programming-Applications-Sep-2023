package jsondata

type Info struct {
	N        int    `json:"n"`
	Min      int    `json:"min"`
	Max      int    `json:"max"`
	Count    int    `json:"count"`
	BasePath string `json:"basePath"`
}

func NewInfo(n, min, max, count int, basePath string) *Info {
	return &Info{N: n, Min: min, Max: max, Count: count, BasePath: basePath}
}
