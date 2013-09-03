package rest

type CorsRequestInfo struct {
	// TODO add a pointer to the rest.Request ?
	IsCors                      bool
	IsPreflight                 bool
	Origin                      string
	AccessControlRequestMethod  string
	AccessControlRequestHeaders string
}

type CorsResponseHeaders struct {
	AccessControlAllowOrigin      string
	AccessControlAllowCredentials string
	AccessControlExposeHeaders    string
	AccessControlMaxAge           string
	AccessControlAllowMethods     string
	AccessControlAllowHeaders     string
}

func newCorsRequestInfo(r *Request) *CorsRequestInfo {

	origin := r.Header.Get("Origin")
	isCors := origin != ""
	reqMethod := r.Header.Get("Access-Control-Request-Method")
	reqHeaders := r.Header.Get("Access-Control-Request-Headers")
	isPreflight := isCors && r.Method == "OPTIONS" && reqMethod != ""

	return &CorsRequestInfo{
		IsCors:      isCors,
		IsPreflight: isPreflight,
		Origin:      origin,
		AccessControlRequestMethod:  reqMethod,
		AccessControlRequestHeaders: reqHeaders,
	}
}

func (self *CorsResponseHeaders) setHeaders(w *ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", self.AccessControlAllowOrigin)
	w.Header().Set("Access-Control-Allow-Credentials", self.AccessControlAllowCredentials)
	w.Header().Set("Access-Control-Expose-Headers", self.AccessControlExposeHeaders)
	w.Header().Set("Access-Control-Max-Age", self.AccessControlMaxAge)
	w.Header().Set("Access-Control-Allow-Methods", self.AccessControlAllowMethods)
	w.Header().Set("Access-Control-Allow-Headers", self.AccessControlAllowHeaders)
}
