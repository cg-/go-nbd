package nbd

import (
	"fmt"
	"os"
  "errors"
)

const (
	logging = true
)

type NbdConnector struct {
	nbd        *NBD
	emunbd     *os.File
	source     *os.File
	mountpoint string
}

func CreateNbdConnector(source, mountpoint string) (*NbdConnector, error) {
  // open the source
	sourceFile, err := os.OpenFile(source, os.O_RDWR, os.FileMode(0777))
  if err != nil {
    return nil, errors.New("couldn't open source")
  }

  // check the destination mountpoint
  _, err = os.Stat(mountpoint)
  if os.IsNotExist(err) {
    return nil, errors.New("couldn't open dest")
  }

	log("created nbd connector with source " + source + " and mountpoint " + mountpoint)
	return &NbdConnector{nil, nil, sourceFile, mountpoint}, nil
}

func (nbdcon *NbdConnector) Dump() string{
  return fmt.Sprintf("nbd: [%s] emunbd: [%s] source: [%s] mountpoint: [%s]", nbdcon.nbd, nbdcon.emunbd, nbdcon.source, nbdcon.mountpoint)
}

func log(msg string) {
	if logging == false {
		return
	}
	fmt.Printf("debug: %s\n", msg)
}
