package jsondata

type Info struct {
	N        int    `json:"n"`
	Min      int    `json:"min"`
	Max      int    `json:"max"`
	Count    int    `json:"count"`
	BasePath string `json:"basePath"`
}
