package lib

import (
	"github.com/taubyte/go-sdk/event"
)

//export getJson
func getJson(e event.Event) uint32 {
	h, err := e.HTTP()
	if err != nil {
		return 1
	}

	h.Headers().Set("Content-Type", "application/json")

	h.Write([]byte(`{"name":"Filecoin Web Services Demo","icon":"https://45fakp4f0.g.aron.lol/fws-small-logo.png","issue":"Helloknfbjklwenfkjwenfkj","styling":{"logo":"https://45fakp4f0.g.aron.lol/fws-full-logo.png","logo-mini":"https://45fakp4f0.g.aron.lol/fws-small-logo.png"}}`))

	return 0
}
