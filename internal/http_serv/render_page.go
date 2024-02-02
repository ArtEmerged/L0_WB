package httpserv

import (
	"html/template"
	"net/http"
	"wblzero/internal/models"

	"github.com/sirupsen/logrus"
)

func (h *Handler) renderPage(w http.ResponseWriter, data *models.Order) {
	temp, err := template.ParseFiles("./ui/templates/index.html")
	if err != nil {
		logrus.Errorf("there was a problem persing the file:%s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = temp.Execute(w, data)
	if err != nil {
		logrus.Errorf("there was a problem executeing the file:%s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
