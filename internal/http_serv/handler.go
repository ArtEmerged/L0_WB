package httpserv

import (
	"net/http"
	"wblzero/internal/models"

	"github.com/sirupsen/logrus"
)

type Order interface {
	Get(uid string) (*models.Order, error)
}

type Handler struct {
	service Order
}

func New(service Order) *Handler {
	return &Handler{service: service}
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		logrus.Errorf("client:[%s] incorrect path:[%s]", r.RemoteAddr, r.URL.Path)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		logrus.Errorf("client:[%s] incorrect method:[%s]", r.RemoteAddr, r.Method)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	h.renderPage(w, nil)
}

func (h *Handler) order(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logrus.Errorf("client:[%s] incorrect method:[%s]", r.RemoteAddr, r.Method)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	orderID := r.FormValue("orderId")

	order, err := h.service.Get(orderID)
	if err != nil {
		logrus.Errorf("geting order:%s%s", err.Error(), orderID)
		if err == models.ErrNoOrder {
			http.Error(w, err.Error()+orderID, http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	h.renderPage(w, order)
}
