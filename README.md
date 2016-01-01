# Important

`go-texttospeech` is only available for Windows for now. Linux will come next and hopefully MAC OSX too :D

# About

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/asticode/go-texttospeech/texttospeech)

`go-texttospeech` is a text to speech manager relying on the OS speech recognition software for the GO programming language (http://golang.org).

# Install `go-texttospeech`

Run the following command:

    $ go get github.com/asticode/go-texttospeech/texttospeech
    
# Example

    // Create text to speech
    tts := texttospeech.NewTextToSpeech()

    // Say
    tts.Say("go-texttospeech is awesome!")
