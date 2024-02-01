package httpserv

import "net/http"

func (h *Handler) InitRouter() http.Handler {
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", h.index)
	mux.HandleFunc("/order/", h.order)
	return mux
}
