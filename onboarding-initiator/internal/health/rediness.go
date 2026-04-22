package health

import "net/http"

type Dependencies struct {

	Kafka bool

	Postgres bool

	Redis bool
}

func Ready(dep Dependencies) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if !dep.Kafka || !dep.Postgres || !dep.Redis {

			http.Error(w,"dependencies not ready",503)

			return
		}

		w.WriteHeader(200)
	}
}
