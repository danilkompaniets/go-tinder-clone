package proxy

import (
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"net/url"
	"strings"
)

func HandleReverseProxy(c *gin.Context, target *url.URL, prefix string) {
	proxy := httputil.NewSingleHostReverseProxy(target)

	r := c.Request.Clone(c.Request.Context())

	r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)
	r.URL.RawPath = strings.TrimPrefix(r.URL.RawPath, prefix)
	r.Host = target.Host

	proxy.ServeHTTP(c.Writer, r)
}
