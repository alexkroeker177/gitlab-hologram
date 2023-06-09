// This file contains the implementation of BlankSparkable.

package app

import (
	"github.com/Bitspark/go-bitnode/bitnode"
	"net/http"
)

// Struct definition for BlankSparkable.

// @@SPARKABLE_DESCRIPTION@@
type BlankSparkable struct {
	bitnode.System

	// Credentials represents the credentials of the Personio account.
	credentials Credentials

	// Custom fields

	// httpClient used for API calls to Personio.
	httpClient *http.Client

	// authToken for the Personio API.
	authToken string
}

// BlankSparkable methods.

// @@METHOD_STUBS@@

// @@HANDLER_STUBS@@

// Lifecycle callbacks.

// lifecycleCreate is called when the container has been created.
func (s *BlankSparkable) lifecycleCreate(vals ...bitnode.HubItem) error {
	// TODO: Add startup logic here which is called when the hologram is created.
	return nil
}

// lifecycleLoad is called when the container has been started (after lifecycleCreate) or restarted.
func (s *BlankSparkable) lifecycleLoad(vals ...bitnode.HubItem) error {
	s.httpClient = &http.Client{}

	s.SetMessage("BlankSparkable running...")
	s.SetStatus(bitnode.SystemStatusRunning)

	return nil
}

// DO NOT CHANGE THE FOLLOWING CODE UNLESS YOU KNOW WHAT YOU ARE DOING.

// Init attaches the methods of the BlankSparkable to the respective handlers.
func (s *BlankSparkable) Init() error {
	// METHODS

	// @@METHOD_HANDLERS@@

	// VALUES

	// @@VALUE_HANDLERS@@

	// CHANNELS

	// @@CHANNEL_HANDLERS@@

	// LIFECYCLE EVENTS

	s.AddCallback(bitnode.LifecycleCreate, bitnode.NewNativeEvent(func(vals ...bitnode.HubItem) error {
		return s.lifecycleCreate(vals...)
	}))

	s.AddCallback(bitnode.LifecycleLoad, bitnode.NewNativeEvent(func(vals ...bitnode.HubItem) error {
		return s.lifecycleLoad(vals...)
	}))

	return nil
}
