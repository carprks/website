package backend

import (
	"fmt"
	"github.com/carprks/website/backend/website"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/keloran/go-probe"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Routes ...
func Routes() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(presetHeaders)

	// Probe
	router.Get("/probe", probe.HTTP)

	// Homepage
	router.Get("/", website.HomeHandler)

	// Privacy
	router.Route("/privacy", func(r chi.Router) {
	  r.Get("/", website.PrivacyHandler)
	  r.Get("/cookie", website.PrivacyCookieHandler)
  })

	// App
	router.Get("/app", website.AppHandler)

	// Companies
	router.Get("/companies", website.CompaniesHandler)

	// About
	router.Get("/about", website.AboutHandler)

	// Contact
	router.Route("/contact", func(r chi.Router) {
		r.Get("/", website.ContactHandler)
		r.Post("/", website.ContactHandler)
	})

	// Pricing
	router.Route("/pricing", func(r chi.Router) {
		r.Get("/", website.PricingHandler)
		r.Post("/", website.PricingHandler)
	})

	// Account
	router.Route("/account", func(r chi.Router) {
		r.Get("/", website.AccountHandler)

    // Login
    r.Route("/login", func(rt chi.Router) {
      rt.Get("/", website.LoginHandler)
      rt.Post("/", website.LoginHandler)
    })

    // Logout
    r.Get("/logout", website.LogoutHandler)

    // Register
    r.Route("/register", func(rt chi.Router) {
      rt.Get("/", website.RegisterHandler)
      rt.Post("/", website.RegisterHandler)
    })
	})

	// CarParks
	router.Route("/carparks", func(r chi.Router) {
		r.Get("/", website.CarParksHandler)
		r.Post("/", website.CarParkHandler)
	})

	// Frontend
	frontEnd(router)

	return router
}

func frontEnd(r chi.Router) {
  distPath := "frontend"
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(fmt.Sprintf("Workdir Err: %v", err))
	}
	paths := []string{
		"css",
		"js",
		"assets",
	}
	for i := 0; i < len(paths); i++ {
		path := paths[i]
		root := filepath.Join(wd, distPath, path)
		fileServer(r, fmt.Sprintf("/%s", path), http.Dir(root))
	}

	css := filepath.Join(wd, distPath, "css")
	js := filepath.Join(wd, distPath, "js")
	fileServer(r, "/css", http.Dir(css))
	fileServer(r, "/js", http.Dir(js))
}

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		fmt.Println(fmt.Sprintf("params not allowed"))
	}

	fs := http.StripPrefix(path, http.FileServer(root))
	if path != "/" && path[len(path) - 1] != '/' {
		r.Get(path, http.RedirectHandler(path + "/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
