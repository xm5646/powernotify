// Copyright 2020 xm5646. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package powernotify

// NotifySender is the interface that the notice message Send method.
//
// Send send Message to receivers and return success number or error.
type NotifySender interface {
	Send() (n int, err error)
}
