package client

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/calmh/randomart"
	charm "github.com/charmbracelet/charm/proto"
	"github.com/charmbracelet/charm/ui/common"
)

var styles = common.DefaultStyles()

// Fingerprint is the fingerprint of an SSH key.
type Fingerprint struct {
	Algorithm string
	Type      string
	Value     string
}

// String outputs a string representation of the fingerprint.
func (f Fingerprint) String() string {
	return fmt.Sprintf(
		"%s %s",
		styles.ListDim.Render(strings.ToUpper(f.Algorithm)),
		styles.ListKey.Render(f.Type+":"+f.Value),
	)
}

// FingerprintSHA256 returns the algorithm and SHA256 fingerprint for the given
// key.
func FingerprintSHA256(k charm.PublicKey) (Fingerprint, error) {
	keyParts := strings.Split(k.Key, " ")
	if len(keyParts) != 2 {
		return Fingerprint{}, charm.ErrMalformedKey
	}

	b, err := base64.StdEncoding.DecodeString(keyParts[1])
	if err != nil {
		return Fingerprint{}, err
	}

	algo := strings.Replace(keyParts[0], "ssh-", "", -1)
	sha256sum := sha256.Sum256(b)
	hash := base64.RawStdEncoding.EncodeToString(sha256sum[:])

	return Fingerprint{
		Algorithm: algo,
		Type:      "SHA256",
		Value:     hash,
	}, nil
}

// RandomArt returns the randomart for the given key.
func RandomArt(k charm.PublicKey) (string, error) {
	keyParts := strings.Split(k.Key, " ")
	if len(keyParts) != 2 {
		return "", charm.ErrMalformedKey
	}

	b, err := base64.StdEncoding.DecodeString(keyParts[1])
	if err != nil {
		return "", err
	}

	algo := strings.ToUpper(strings.Replace(keyParts[0], "ssh-", "", -1))

	// TODO: also add bit size of key
	h := sha256.New()
	_, _ = h.Write(b)
	board := randomart.GenerateSubtitled(h.Sum(nil), algo, "SHA256").String()
	return strings.TrimSpace(board), nil
}
