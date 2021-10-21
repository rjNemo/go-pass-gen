package api

func (s Server) routes() {
	s.Router.Post("/new", s.HandleNewPassword)
}
