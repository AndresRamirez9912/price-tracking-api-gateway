package handlers

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"price-tracking-api-gateway/src/constants"
	"price-tracking-api-gateway/src/services"
)

func ForwardingV1(w http.ResponseWriter, r *http.Request) {
	targetURL, err := services.GetTargetRoute(r)
	if err != nil {
		return
	}
	// Add the User from the context to the request Body
	user := r.Context().Value(constants.USER_CONTEXT)
	if user != nil {
		err = services.AddBodyElement(r, constants.USER_CONTEXT, user)
		if err != nil {
			return
		}
	}

	// Forward the Request to the target
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.ErrorLog = log.New(os.Stderr, constants.PROXY_ERROR, log.LstdFlags)
	proxy.Director = func(request *http.Request) {
		// Update the Path of the new req
		request.URL.Scheme = targetURL.Scheme
		request.URL.Host = targetURL.Host
		request.URL.Path = r.URL.Path
		request.Host = request.URL.Host

		// Add Headers
		services.AddHeaders(request)
	}

	// Forward Request
	proxy.ServeHTTP(w, r)

}
