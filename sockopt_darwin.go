// Copyright 2014 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tcp

import (
	"os"
	"syscall"
	"time"
)

func (c *Conn) setKeepAliveIdlePeriod(d time.Duration) error {
	fd, err := c.sysfd()
	if err != nil {
		return err
	}
	d += (time.Second - time.Nanosecond)
	secs := int(d.Seconds())
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd, ianaProtocolTCP, sysTCP_KEEPALIVE, secs))
}

func (c *Conn) setKeepAliveProbeInterval(d time.Duration) error {
	fd, err := c.sysfd()
	if err != nil {
		return err
	}
	d += (time.Second - time.Nanosecond)
	secs := int(d.Seconds())
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd, ianaProtocolTCP, sysTCP_KEEPINTVL, secs))
}

func (c *Conn) setMaxKeepAliveProbes(max int) error {
	fd, err := c.sysfd()
	if err != nil {
		return err
	}
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd, ianaProtocolTCP, sysTCP_KEEPCNT, max))
}

func (c *Conn) setCork(on bool) error {
	fd, err := c.sysfd()
	if err != nil {
		return err
	}
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd, ianaProtocolTCP, sysTCP_NOPUSH, boolint(on)))
}

func (c *Conn) info() (*Info, error) {
	return nil, errOpNoSupport
}
