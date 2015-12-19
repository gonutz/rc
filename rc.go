package rc

import (
	"encoding/binary"
	"os"
	"time"
)

type Key uint32

const (
	KeyOff            Key = 2
	KeyTVRecord           = 24
	KeyTVText             = 49
	KeyTVPlay             = 28
	KeyTVEye              = 4
	KeyRed                = 50
	KeyGreen              = 51
	KeyYellow             = 52
	KeyBlue               = 53
	KeyMute               = 0
	KeyTVT                = 22
	KeyVolumeUp           = 9
	KeyVolumeDown         = 8
	KeyProgramUp          = 11
	KeyProgramDown        = 12
	KeyLeft               = 29
	KeyRight              = 31
	KeyUp                 = 26
	KeyDown               = 34
	KeyOK                 = 30
	KeyBack               = 32
	KeyWindows            = 27
	KeyInfo               = 47
	KeyChapterBack        = 33
	KeyChapterForward     = 35
	KeyRecord             = 39
	KeyPause              = 41
	KeyStop               = 40
	KeyRewind             = 36
	KeyFastForward        = 38
	KeyPlay               = 37
	Key1                  = 13
	Key2                  = 14
	Key3                  = 15
	Key4                  = 16
	Key5                  = 17
	Key6                  = 18
	Key7                  = 19
	Key8                  = 20
	Key9                  = 21
	Key0                  = 23
	KeyAsterisk           = 55
	KeyHash               = 56
	KeyClear              = 48
	KeyEnter              = 54
)

// OpenInput tries to open /dev/input/event0 and starts reading keys from
// the remote control in an endless loop. If the file cannot be opened or
// if there is an error reading it (in case the remote control was
// removed) it will try again after sleeping a second. This happens in an
// infinite loop.
func OpenInput() <-chan Key {
	c := make(chan Key)
	go func() {
		for {
			input, err := os.Open("/dev/input/event0")
			if err != nil {
				time.Sleep(time.Second)
				continue
			}
			for {
				var event event
				err = binary.Read(input, binary.LittleEndian, &event)
				if err != nil {
					break
				}
				if event.Type == miscEvent {
					c <- Key(event.Value)
				}
			}
			input.Close()
		}
	}()
	return c
}

type event struct {
	TimeStamp uint64
	Type      uint16
	Code      uint16
	Value     uint32
}

const miscEvent = 0x4
