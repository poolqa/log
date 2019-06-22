// +build !linux

package rotator

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error {
	return nil
}
