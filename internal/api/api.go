package api

import (
	"github.com/garieiro/ask-me-something-backend.git/internal/store/pgstore"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/gorilla/websocket"
	"net/http"
)

type apiHandler struct {
	q        *pgstore.Queries //TODO: SHOULD BE A INTERFACE
	r        *chi.Mux
	upgrader websocket.Upgrader
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	a := apiHandler{q: q, upgrader: websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}}

	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		//Atention!!! This is not correct but as it is to be run locally, there is not much of a problem.
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/subscribe/{room_id}", a.handleSubscribe)

	r.Route("/api", func(r chi.Router) {
		r.Route("/rooms", func(r chi.Router) {
			r.Post("/", a.handleCreateRoom)
			r.Get("/", a.handleGetRooms)

			r.Route("/{room_id}/messages", func(r chi.Router) {
				r.Get("/", a.handleCreateRoomMessage)
				r.Post("/", a.handleCreateMessages)

				r.Route("/{message_id}", func(r chi.Router) {
					r.Get("/", a.handleGetRoomMessage)
					r.Patch("/react", a.handleReactToMessage)
					r.Delete("/react", a.handleRemoveReactFromMessage)
					r.Patch("/answer", a.handleMarkAsAnsweredFromMessage)
				})
			})
		})
	})

	a.r = r
	return a
}

func (h apiHandler) handleSubscribe(w http.ResponseWriter, r *http.Request)                 {}
func (h apiHandler) handleCreateRoom(w http.ResponseWriter, r *http.Request)                {}
func (h apiHandler) handleGetRooms(w http.ResponseWriter, r *http.Request)                  {}
func (h apiHandler) handleCreateRoomMessage(w http.ResponseWriter, r *http.Request)         {}
func (h apiHandler) handleCreateMessages(w http.ResponseWriter, r *http.Request)            {}
func (h apiHandler) handleGetRoomMessage(w http.ResponseWriter, r *http.Request)            {}
func (h apiHandler) handleReactToMessage(w http.ResponseWriter, r *http.Request)            {}
func (h apiHandler) handleRemoveReactFromMessage(w http.ResponseWriter, r *http.Request)    {}
func (h apiHandler) handleMarkAsAnsweredFromMessage(w http.ResponseWriter, r *http.Request) {}
