package texttospeech

import (
	"os/exec"
)

func (tts textToSpeech) say(i string) error {

	command := exec.Command("say", "\""+i+"\"")
	stdoutStderr, e := command.CombinedOutput()

	// Process error
	if e != nil {
		return e
	}
}
