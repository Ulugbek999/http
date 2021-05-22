package app

import (
	"github.com/Ulugbek999/http.git/pkg/types"
	"github.com/Ulugbek999/http.git/pkg/banners"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	
)

//Server - represents the logical server of application.
type Server struct {
	mux        *http.ServeMux
	bannersSvc *banners.Service
}

//NewServer - constructor function to create a new server.
func NewServer(mux *http.ServeMux, bannersSvc *banners.Service) *Server {
	return &Server{
		mux:        mux,
		bannersSvc: bannersSvc,
	}
}

//ServeHTTP - method to start the server.
func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.mux.ServeHTTP(writer, request)
}

//Init - initializes the server (register all handlers).
func (s *Server) Init() {
	s.mux.HandleFunc("/banners.getAll", s.handleGetAllBanners)
	s.mux.HandleFunc("/banners.getById", s.handleGetBannerByID)
	s.mux.HandleFunc("/banners.save", s.handleSaveBanner)
	s.mux.HandleFunc("/banners.removeById", s.handleRemoveByID)
}

//handleGetAllBanners - handler method for extracting all banners.
func (s *Server) handleGetAllBanners(writer http.ResponseWriter, request *http.Request) {

	items, err := s.bannersSvc.All(request.Context())
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(items)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	responseByJSON(writer, data)
}

//handleGetbannerByID - handler method for extracting banner by id.
func (s *Server) handleGetBannerByID(writer http.ResponseWriter, request *http.Request) {
	idParam := request.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	item, err := s.bannersSvc.ByID(request.Context(), id)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(item)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	responseByJSON(writer, data)
}

//handleSaveBanner - handler method for saves/updates banner.
func (s *Server) handleSaveBanner(writer http.ResponseWriter, request *http.Request) {

	idParam := request.URL.Query().Get("id")
	title := request.URL.Query().Get("title")
	content := request.URL.Query().Get("content")
	button := request.URL.Query().Get("button")
	link := request.URL.Query().Get("link")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if title == "" || content == "" || button == "" || link == "" {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	items := &types.Banner{
		ID:      id,
		Title:   title,
		Content: content,
		Button:  button,
		Link:    link,
	}

	banner, err := s.bannersSvc.Save(request.Context(), items)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(banner)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	responseByJSON(writer, data)
}

//handleRemoveByID - handler method for removes banner by ID.
func (s *Server) handleRemoveByID(writer http.ResponseWriter, request *http.Request) {
	idParam := request.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	item, err := s.bannersSvc.RemoveByID(request.Context(), id)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(item)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	responseByJSON(writer, data)
}

//responseByJSON - response from JSON.
func responseByJSON(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(data)
	if err != nil {
		log.Println(err)
	}
}
