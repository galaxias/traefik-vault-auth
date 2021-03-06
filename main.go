package traefik_vault_auth

import (
    "fmt"
	"context"
	"net/http"
)

// VaultAuth a plugin to use Vault as authentication provider for Basic Auth Traefik middleware.
type VaultAuth struct {
	next   http.Handler
	name   string
	config *Config
}

// New created a new VaultBasicAuth plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &VaultAuth{
		next:   next,
		name:   name,
		config: config,
	}, nil
}

func (va *VaultAuth) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	user, pass, ok := req.BasicAuth()

	if !ok {
		// No valid 'Authentication: Basic xxxx' header found in request
		rw.Header().Set("WWW-Authenticate", `Basic realm="`+va.config.CustomRealm+`"`)
		http.Error(rw, "Unauthorized.", http.StatusUnauthorized)
		return
	}
	fmt.Println("Begin to execute vault login with user " + user + " password " + pass)
	if err := va.config.Vault.login(user, pass); err != nil {
		// Failed to login with provided user/pass
		rw.Header().Set("WWW-Authenticate", `Basic realm="`+va.config.CustomRealm+`"`)
		http.Error(rw, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	va.next.ServeHTTP(rw, req)
}
