// SPDX-License-Identifier: GPL-3.0-or-later

package samba

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/netdata/netdata/go/plugins/pkg/executable"
)

func (s *Samba) initSmbStatusBinary() (smbStatusBinary, error) {
	ndsudoPath := filepath.Join(executable.Directory, "ndsudo")
	if _, err := os.Stat(ndsudoPath); err != nil {
		return nil, fmt.Errorf("ndsudo executable not found: %v", err)

	}

	smbStatus := newSmbStatusBinary(ndsudoPath, s.Timeout.Duration(), s.Logger)

	return smbStatus, nil
}
