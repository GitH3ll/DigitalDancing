package dance

import "os"

type Dancer interface {
	Dance(file *os.File)
}
