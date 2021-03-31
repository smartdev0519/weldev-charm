package client

import (
	"encoding/json"

	charm "github.com/charmbracelet/charm/proto"
	"github.com/dgrijalva/jwt-go"
)

// Auth is the authenticated user's charm id and jwt returned from the ssh server.
type Auth struct {
	CharmID     string              `json:"charm_id"`
	JWT         string              `json:"jwt"`
	PublicKey   string              `json:"public_key"`
	EncryptKeys []*charm.EncryptKey `json:"encrypt_keys"`
	claims      *jwt.StandardClaims
}

// Auth returns the Auth struct for a client session. It will renew and cache
// the Charm ID JWT.
func (cc *Client) Auth() (*Auth, error) {
	cc.authLock.Lock()
	defer cc.authLock.Unlock()

	if cc.auth.claims == nil || cc.auth.claims.Valid() != nil {
		auth := &Auth{}
		s, err := cc.sshSession()
		if err != nil {
			return nil, charm.ErrAuthFailed{Err: err}
		}
		defer s.Close()

		b, err := s.Output("api-auth")
		if err != nil {
			return nil, charm.ErrAuthFailed{Err: err}
		}
		err = json.Unmarshal(b, auth)
		if err != nil {
			return nil, charm.ErrAuthFailed{Err: err}
		}

		p := &jwt.Parser{}
		token, _, err := p.ParseUnverified(auth.JWT, &jwt.StandardClaims{})
		if err != nil {
			return nil, charm.ErrAuthFailed{Err: err}
		}
		auth.claims = token.Claims.(*jwt.StandardClaims)
		cc.auth = auth
		if err != nil {
			return nil, charm.ErrAuthFailed{Err: err}
		}
	}
	return cc.auth, nil
}

// InvalidateAuth clears the JWT auth cache, forcing subsequent Auth() to fetch
// a new JWT from the server.
func (cc *Client) InvalidateAuth() {
	cc.authLock.Lock()
	defer cc.authLock.Unlock()
	cc.auth.claims = nil
}
