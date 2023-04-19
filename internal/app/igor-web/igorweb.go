// Copyright 2023 National Technology & Engineering Solutions of Sandia, LLC (NTESS).
// Under the terms of Contract DE-NA0003525 with NTESS, the U.S. Government retains
// certain rights in this software.

package igorweb

import (
	"time"
)

var igorweb struct {
	Config     // embed
	ConfigPath string
	IgorHome   string // embed
	Started    time.Time
}
