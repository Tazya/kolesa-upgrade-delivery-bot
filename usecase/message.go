package usecase

type Message struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Sender interface {
	SendAll(msg Message) error
}
