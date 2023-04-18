package api

import (
	"encoding/json"
	"fmt"
	"github.com/TelenLiu/ip/models"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"strings"
)

func GetRealIP(r *http.Request) (ip string) {
	ip = net.ParseIP(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0]).String()
	if ip == "<nil>" {
		ip = net.ParseIP(strings.Split(r.Header.Get("X-Real-Ip"), ",")[0]).String()
	}
	if ip == "<nil>" {
		remoteAddr, _, _ := net.SplitHostPort(r.RemoteAddr)
		if remoteAddr == "::1" {
			ip = "127.0.0.1"
		}
	}
	return
}

// GetIP returns a user's public facing IP address (IPv4 OR IPv6).
//
// By default, it will return the IP address in plain text, but can also return
// data in both JSON and JSONP if requested to.
func GetIP(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// We'll always grab the first IP address in the X-Forwarded-For header
	// list.  We do this because this is always the *origin* IP address, which
	// is the *true* IP of the user.  For more information on this, see the
	// Wikipedia page: https://en.wikipedia.org/wiki/X-Forwarded-For
	ip := GetRealIP(r)

	// If the user specifies a 'format' querystring, we'll try to return the
	// user's IP address in the specified format.
	if format, ok := r.Form["format"]; ok && len(format) > 0 {
		jsonStr, _ := json.Marshal(models.IPAddress{ip})

		switch format[0] {
		case "json":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(jsonStr))
			return
		case "jsonp":
			// If the user specifies a 'callback' parameter, we'll use that as
			// the name of our JSONP callback.
			callback := "callback"
			if val, ok := r.Form["callback"]; ok && len(val) > 0 {
				callback = val[0]
			}

			w.Header().Set("Content-Type", "application/javascript")
			fmt.Fprintf(w, callback+"("+string(jsonStr)+");")
			return
		}
	}

	// If no 'format' querystring was specified, we'll default to returning the
	// IP in plain text.
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, ip)
}
