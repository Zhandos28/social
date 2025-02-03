package main

import (
	"context"
	"errors"
	"github.com/Zhandos28/social/internal/store"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type userKey string

const (
	userCtxKey userKey = "user"
)

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := app.getUserContext(r)

	jsonResponse(w, http.StatusOK, user)
	return
}

type FollowUser struct {
	UserID int64 `json:"user_id"`
}

func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	follower := app.getUserContext(r)

	// Revert back to auth userID from ctx
	var payload FollowUser
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if err := app.store.Followers.Follow(r.Context(), payload.UserID, follower.ID); err != nil {
		switch err {
		case store.ErrConflict:
			app.conflictError(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	jsonResponse(w, http.StatusNoContent, payload)
}

func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	unfollower := app.getUserContext(r)

	// Revert back to auth userID from ctx
	var payload FollowUser
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if err := app.store.Followers.Unfollow(r.Context(), payload.UserID, unfollower.ID); err != nil {
		switch err {
		case store.ErrConflict:
			app.conflictError(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	jsonResponse(w, http.StatusNoContent, payload)
}

func (app *application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
		if err != nil {
			app.badRequestError(w, r, err)
			return
		}

		user, err := app.store.Users.GetByID(r.Context(), userID)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundError(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
		}

		ctx := context.WithValue(r.Context(), userCtxKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) getUserContext(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtxKey).(*store.User)
	return user
}
