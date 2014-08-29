package gopbf

import (
    "github.com/edsrzf/mmap-go"
)

func New(fileName string) {
    file, err := os.Open(fileName)
    pError(err)

    buffer, err := mmap.Map(file, mmap.RDONLY, 0)
    pError(err)
}
