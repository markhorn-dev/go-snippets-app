package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/markhorn-dev/go-snippets-app/internal/models"
)

// ===============================================================================
// home is the handler for /
// ===============================================================================

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, http.StatusOK, "home.html", data)
}

// ===============================================================================
// snippetView is the handler for /snippet/view/123
// ===============================================================================

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	sid := chi.URLParamFromCtx(ctx, "id")
	id, err := strconv.Atoi(sid)

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, http.StatusOK, "view.html", data)
}

// ===============================================================================
// snippetCreate is the get handler for /snippet/create
// ===============================================================================

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create/Get"))
}

// ===============================================================================
// snippetCreate is the post handler for /snippet/create
// ===============================================================================

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create/Post"))
}
