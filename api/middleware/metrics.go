package middleware

import (
	"net/http"
	"strconv"

	"github.com/martinyonatann/go-invoice/pkg/metric"
	"github.com/urfave/negroni"
)

func Metrics(mService metric.Service) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, handler http.HandlerFunc) {
		appMetric := metric.NewHTTP(r.URL.Path, r.Method)
		appMetric.StartedHttp()

		handler(w, r)

		res := w.(negroni.ResponseWriter)

		appMetric.FinishedHttp()
		appMetric.StatusCode = strconv.Itoa(res.Status())

		mService.SaveHTTP(appMetric)
	}
}
