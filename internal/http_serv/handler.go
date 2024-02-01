package httpserv

import (
	"fmt"
	"html/template"
	"net/http"
	"wblzero/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	temp, err := template.ParseFiles("./ui/templates/index.html")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	err = temp.Execute(w, nil)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}

func (h *Handler) order(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/order/" {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	orderID := r.FormValue("orderId")
	fmt.Println("order" + orderID)
	_, err := h.service.Get(orderID)
	if err != nil {
		fmt.Printf("errGet:%s", err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	temp, err := template.ParseFiles("./ui/templates/index.html")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	err = temp.Execute(w, nil)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}
