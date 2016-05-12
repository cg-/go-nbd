package nbd

import (
	"fmt"
	"os"
  "errors"
  "time"
  "syscall"
  "path/filepath"
)

const (
	logging = true
)

type NbdConnector struct {
	nbd        *NBD
	source     *os.File
	emunbd     string
	mountpoint string
}

func CreateNbdConnector(source, mountpoint string) (*NbdConnector, error) {
  // open the source
	sourceFile, err := os.OpenFile(source, os.O_RDWR, os.FileMode(0777))
  if err != nil {
    return nil, errors.New("couldn't open source file")
  }

  // check the destination mountpoint
  _, err = os.Stat(mountpoint)
  if os.IsNotExist(err) {
    return nil, errors.New("mountpoint doesn't exist")
  }

	stat, _ := sourceFile.Stat()
	dev := Create(sourceFile, stat.Size())

  go dev.Connect()
  time.Sleep(1 * time.Second)

	emu := dev.GetName()
  mountpointAbs, err := filepath.Abs(mountpoint)

	return &NbdConnector{dev, sourceFile, emu, mountpointAbs}, nil
}

func (nbdcon *NbdConnector) Mount() error {
  err := syscall.Mount(nbdcon.emunbd, nbdcon.mountpoint, "ext3", 0, "") != nil
	if err == true {
		return errors.New("Couldn't mount")
	}
	return nil
}

func (nbdcon *NbdConnector) Unmount() error {
  err := syscall.Unmount(nbdcon.mountpoint, 0) != nil
	if err == true {
		return errors.New("Couldn't unmount")
	}
	return nil
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
