package gokeyboard

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"
	"os"
	"syscall"
	"unsafe"
)

const (
	keyEventSizeInBytes = unsafe.Sizeof(Key{})
)

type KeyboardFile struct {
	fd *os.File
}

func (k *KeyboardFile) Close() error {
	if k.fd == nil {
		return nil
	}
	return k.fd.Close()
}

// Will return a stream of Keys.
func (k *KeyboardFile) Read() chan Key {
	event := make(chan Key)

	go func(eventChan chan Key) {
		for {
			event, err := k.readKey()
			if err != nil {
				close(eventChan)
				log.Fatal(err)
			}
			if event != nil && event.isKeyEvent() {
				eventChan <- *event
			}

		}
	}(event)

	return event
}

// Will read a buffer from the keyboard file, and will try to export it to key
func (k *KeyboardFile) readKey() (*Key, error) {
	buffer := make([]byte, keyEventSizeInBytes)
	bytesRead, err := k.fd.Read(buffer)
	if err != nil {
		return nil, err
	}
	if bytesRead <= 0 {
		return nil, nil
	}

	return k.generateKeyFromBuffer(buffer)
}

func (k *KeyboardFile) generateKeyFromBuffer(buffer []byte) (*Key, error) {
	event := &Key{}
	err := binary.Read(bytes.NewReader(buffer), binary.LittleEndian, event)
	return event, err
}

// Will check if the program runs as root
func IsRoot() bool {
	return syscall.Getuid() == 0 && syscall.Geteuid() == 0
}

// Will create a new Keyboard listener
func New(keyboardFilePath string) (*KeyboardFile, error) {
	keyboard := &KeyboardFile{}
	if !IsRoot() {
		return nil, errors.New("Must run as root to read keyboard file.")
	}
	fd, err := os.Open(keyboardFilePath)
	keyboard.fd = fd
	return keyboard, err
}
