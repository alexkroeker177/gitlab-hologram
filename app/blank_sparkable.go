// This file contains the implementation of BlankSparkable.

package app

import (
	"github.com/Bitspark/go-bitnode/bitnode"
)

// Struct definition for BlankSparkable.

// BlankSparkable is the main sparkable.
// @@SPARKABLE_DESCRIPTION@@
type BlankSparkable struct {
	bitnode.System

	// @@SPARKABLE_FIELDS@@
}

// BlankSparkable methods.

// @@METHOD_STUBS@@

// @@HANDLER_STUBS@@

// Lifecycle callbacks.

// lifecycleCreate is called when the container has been created.
func (s *BlankSparkable) lifecycleCreate(vals ...bitnode.HubItem) error {
	// TODO: Add startup logic here which is called when the spark is created.
	return nil
}

// lifecycleLoad is called when the container has been started (after lifecycleCreate) or restarted.
func (s *BlankSparkable) lifecycleLoad(vals ...bitnode.HubItem) error {
	// TODO: Add startup logic here which is called after the spark has been created.

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
