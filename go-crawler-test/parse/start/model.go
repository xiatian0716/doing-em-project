package start

type Start struct {
	StartUrl string
}

func (t Start) String() string {
	return "StartUrl:" + t.StartUrl
}
