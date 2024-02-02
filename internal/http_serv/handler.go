package httpserv

import (
	"net/http"
	"wblzero/internal/service"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/order/" {
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
	if r.URL.Path != "/order/" {
		logrus.Errorf("client:[%s] incorrect path:[%s]", r.RemoteAddr, r.URL.Path)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		logrus.Errorf("client:[%s] incorrect method:[%s]", r.RemoteAddr, r.Method)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	orderID := r.FormValue("orderId")

	order, err := h.service.Get(orderID)
	if err != nil {
		logrus.Errorf("errGet:%s", err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
	h.renderPage(w, order)
}
