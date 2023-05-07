package service

import (
	"log"
	"net/http"
	"text/template"
)

type ViewData struct {
	Order string
}

func (s *Service) GetByUUID(w http.ResponseWriter, r *http.Request) {
	uuid, data := r.URL.Query().Get("uuid"), ViewData{Order: ""}
	if uuid != "" {
		order := s.store.GetByUUID(uuid)
		if order == "" {
			order = "Не найден order"
		}
		data.Order = order
	}

	ts, err := template.ParseFiles("internal/service/static/index.tmpl")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	if err = ts.Execute(w, data); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
