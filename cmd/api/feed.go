package main

import (
	"github.com/Zhandos28/social/internal/store"
	"net/http"
)

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
		Tags:   []string{},
		Search: "",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	ctx := r.Context()

	posts, err := app.store.Posts.GetUserFeed(ctx, int64(67), fq)
	if err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if err := jsonResponse(w, http.StatusOK, posts); err != nil {
		app.badRequestError(w, r, err)
		return
	}
}
