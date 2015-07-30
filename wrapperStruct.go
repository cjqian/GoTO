package outputFormatter

func MakeWrapper(r interface{}) Wrapper {
	w := Wrapper{r, 1.1}
	return w
}

type Wrapper struct {
	Resp    interface{} `json:"response"`
	Version float64     `json:"version"`
}
