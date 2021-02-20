package controllers

import "github.com/victorsteven/fullstack/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")

	//Tests
	s.Router.HandleFunc("/tests", middlewares.SetMiddlewareJSON(s.GetTests)).Methods("GET")
	s.Router.HandleFunc("/tests2", middlewares.SetMiddlewareJSON(s.GetTestsAndVars)).Methods("GET")
	s.Router.HandleFunc("/testsvars", middlewares.SetMiddlewareJSON(s.GetTestsVars)).Methods("GET")
	s.Router.HandleFunc("/test", middlewares.SetMiddlewareJSON(s.GetTest)).Methods("GET")
	s.Router.HandleFunc("/GetKattaTest", middlewares.SetMiddlewareJSON(s.GetKattaTest)).Methods("GET")

	//register user
	s.Router.HandleFunc("/register", middlewares.SetMiddlewareJSON(s.Register)).Methods("POST")

}
