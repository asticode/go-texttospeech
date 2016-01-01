package texttospeech

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func (tts textToSpeech) say(i string) error {
	// Init COM
	e := ole.CoInitialize(0)
	defer ole.CoUninitialize()

	// Process error
	if e != nil {
		return e
	}

	// Create object
	unknown, e := oleutil.CreateObject("SAPI.SpVoice")

	// Process error
	if e != nil {
		return e
	}

	// Get voice
	voice, e := unknown.QueryInterface(ole.IID_IDispatch)

	// Process error
	if e != nil {
		return e
	}

	// Speak
	_, e = oleutil.CallMethod(voice, "Speak", i)

	// Return
	return e
}
