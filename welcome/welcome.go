package welcome

import (
	"log"
	"net/http"
)

type Welcome struct {
	log *log.Logger
}

func NewWelcome(l *log.Logger) *Welcome {
	return &Welcome{
		log: l,
	}
}

func (w *Welcome) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	w.log.Println("Welcome to WELCOMEPAGE")
}
