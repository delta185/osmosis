package authenticator

import (
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/osmosis-labs/osmosis/v19/x/authenticator/iface"
)

type AuthenticatorManager struct {
	registeredAuthenticators  []iface.Authenticator
	defaultAuthenticatorIndex int
}

// NewAuthenticatorManager creates a new AuthenticatorManager.
func NewAuthenticatorManager() *AuthenticatorManager {
	return &AuthenticatorManager{
		registeredAuthenticators:  []iface.Authenticator{},
		defaultAuthenticatorIndex: -1,
	}
}

// ResetAuthenticators resets all registered authenticators.
func (am *AuthenticatorManager) ResetAuthenticators() {
	am.registeredAuthenticators = []iface.Authenticator{}
}

// InitializeAuthenticators initializes authenticators. If already initialized, it will not overwrite.
func (am *AuthenticatorManager) InitializeAuthenticators(initialAuthenticators []iface.Authenticator) {
	if len(am.registeredAuthenticators) > 0 {
		return
	}
	am.registeredAuthenticators = initialAuthenticators
}

// RegisterAuthenticator adds a new authenticator to the list of registered authenticators.
func (am *AuthenticatorManager) RegisterAuthenticator(authenticator iface.Authenticator) {
	am.registeredAuthenticators = append(am.registeredAuthenticators, authenticator)
}

// UnregisterAuthenticator removes an authenticator from the list of registered authenticators.
func (am *AuthenticatorManager) UnregisterAuthenticator(authenticator iface.Authenticator) {
	for i, auth := range am.registeredAuthenticators {
		if auth == authenticator {
			am.registeredAuthenticators = append(am.registeredAuthenticators[:i], am.registeredAuthenticators[i+1:]...)
			break
		}
	}
}

// GetRegisteredAuthenticators returns the list of registered authenticators.
func (am *AuthenticatorManager) GetRegisteredAuthenticators() []iface.Authenticator {
	return am.registeredAuthenticators
}

// IsAuthenticatorTypeRegistered checks if the authenticator type is registered.
func (am *AuthenticatorManager) IsAuthenticatorTypeRegistered(authenticatorType string) bool {
	for _, authenticator := range am.GetRegisteredAuthenticators() {
		if authenticator.Type() == authenticatorType {
			return true
		}
	}
	return false
}

// GetAuthenticatorByType returns the base implementation of the authenticator type
func (am *AuthenticatorManager) GetAuthenticatorByType(authenticatorType string) iface.Authenticator {
	for _, authenticator := range am.GetRegisteredAuthenticators() {
		if authenticator.Type() == authenticatorType {
			return authenticator
		}
	}
	return nil
}

// SetDefaultAuthenticatorIndex sets the default authenticator index.
func (am *AuthenticatorManager) SetDefaultAuthenticatorIndex(index int) {
	am.defaultAuthenticatorIndex = index
	if am.defaultAuthenticatorIndex < 0 || am.defaultAuthenticatorIndex >= len(am.registeredAuthenticators) {
		panic("Invalid default authenticator index")
	}
}

// GetDefaultAuthenticator retrieves the default authenticator.
func (am *AuthenticatorManager) GetDefaultAuthenticator() iface.Authenticator {
	if am.defaultAuthenticatorIndex < 0 {
		// ToDo: Instead of panicking, maybe return a FalseAuthenticator that never authenticates?
		panic("Default authenticator not set")
	}
	return am.registeredAuthenticators[am.defaultAuthenticatorIndex]
}

type TransientStore struct {
	storeKey     sdk.StoreKey
	transientCtx sdk.Context
}

func NewTransientStore(storeKey sdk.StoreKey, ctx sdk.Context) *TransientStore {
	return &TransientStore{
		storeKey:     storeKey,
		transientCtx: ctx,
	}
}

func (as *TransientStore) ResetTransientContext(ctx sdk.Context) sdk.Context {
	as.transientCtx, _ = ctx.CacheContext()
	return as.transientCtx
}

func (as *TransientStore) GetKvStore() store.KVStore {
	return as.transientCtx.KVStore(as.storeKey)
}

func (as *TransientStore) Write(ctx sdk.Context) {
	if as.transientCtx.IsZero() {
		panic("Transient context not set")
	}

	// TODO is there a better way to do this? Ideally we would just copy/move the entire store
	ctxStore := ctx.KVStore(as.storeKey)
	iter := as.transientCtx.KVStore(as.storeKey).Iterator(nil, nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		ctxStore.Set(iter.Key(), iter.Value())
	}
}