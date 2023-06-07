package main

import "github.com/Bitspark/go-bitnode/bitnode"

// BlankSparkable is a Bitnode sparkable (more info at https://bitnode.bitspark.com).
//
// /* @BITNODE_DESCRIPTION@ */
//
// URL: https://bitspark.studio/sparkable/blankDomain.blankSparkable
type BlankSparkable struct {
	bitnode.System
}

// Lifecycle callbacks.

/* @BITNODE_METHODS@ */

// lifecycleCreate is called when the container has been created.
func (s *BlankSparkable) lifecycleCreate(vals ...bitnode.HubItem) error {
	s.SetMessage("BlankSparkable running...")
	s.SetStatus(bitnode.SystemStatusRunning)
	return nil
}

// lifecycleLoad is called when the container has been started (after lifecycleCreate) or restarted.
func (s *BlankSparkable) lifecycleLoad(vals ...bitnode.HubItem) error {
	s.SetMessage("BlankSparkable running...")
	s.SetStatus(bitnode.SystemStatusRunning)
	return nil
}

// Init method.

// init attaches the methods of this BlankSparkable to the respective handlers.
func (s *BlankSparkable) init() error {
	// METHODS

	/* @BITNODE_METHOD_CBS@ */

	// LIFECYCLE EVENTS

	s.AddCallback(bitnode.LifecycleCreate, bitnode.NewNativeEvent(func(vals ...bitnode.HubItem) error {
		return s.lifecycleCreate(vals...)
	}))

	s.AddCallback(bitnode.LifecycleLoad, bitnode.NewNativeEvent(func(vals ...bitnode.HubItem) error {
		return s.lifecycleLoad(vals...)
	}))

	return nil
}
