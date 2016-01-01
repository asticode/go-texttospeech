package texttospeech

// TextToSpeech represents an object capable of saying things
type TextToSpeech interface {
	Say(i string) error
}

// NewTextToSpeech creates a new TextToSpeech
func NewTextToSpeech() TextToSpeech {
	return &textToSpeech{}
}

type textToSpeech struct {}

func (tts textToSpeech) Say(i string) error {
	return tts.say(i)
}
