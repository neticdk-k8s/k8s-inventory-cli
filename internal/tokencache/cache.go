package tokencache

import (
	"github.com/neticdk-k8s/k8s-inventory-cli/internal/oidc"
)

// Cache represents a token caching interface
type Cache interface {
	Lookup(key Key) (*oidc.TokenSet, error)
	Save(key Key, tokenSet oidc.TokenSet) error
	Invalidate(key Key) error
}
