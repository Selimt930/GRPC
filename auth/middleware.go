package auth

import "net/http"

type Verifier struct {
	verification func(r *http.Request) (*http.Request, error)
}

func NewReqVerification(v func(r *http.Request) (*http.Request, error)) Verifier {
	return Verifier{verification: v}
}

func (v Verifier) RequestVerification(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := v.verification(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		handler(w, req)
	}
}
