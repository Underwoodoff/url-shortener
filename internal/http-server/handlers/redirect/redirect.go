package redirect

import (
	"errors"
	"log/slog"
	"net/http"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// URLGetter определяет интерфейс для получения URL по алиасу.
type URLGetter interface {
	GetURL(alias string) (string, error)
}

func New(log *slog.Logger, urlGetter URLGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "heandlers.url.redirect.New"

		log = log.With(
			slog.String("op", op),
			slog.String("requst_id", middleware.GetReqID(r.Context())),
		)
		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Error("alias is empty")
			http.Error(w, "alias is required", http.StatusBadRequest)
			return
		}

		originalURL, err := urlGetter.GetURL(alias)
		if errors.Is(err, storage.ErrAliasNotFound) {
			log.Warn("alias not found", slog.String("alias", alias))
			http.Error(w, "alias not found", http.StatusNotFound)
			return
		}

		if err != nil {
			log.Error("failed to get URL", sl.Err(err))
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		log.Info("redirecting", slog.String("alias", alias), slog.String("url", originalURL))
		http.Redirect(w, r, originalURL, http.StatusFound) // 302 Redirect
	}

}
