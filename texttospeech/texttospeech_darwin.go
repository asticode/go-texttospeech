package texttospeech

import (
	"os/exec"
)

func (tts textToSpeech) say(i string) error {

	command := exec.Command("sh", "-c", "say \""+i+"\"")
	err := command.Start()

	if err != nil {
		panic(err)
	}
}
