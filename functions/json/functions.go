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

	h.Write([]byte(`{"name":"Filecoin Web Services Demo","icon":"https://b3tp9jag0.g.tfil.cloud/fws-small-logo.png","issue":"https://b3tp9jag0.g.tfil.cloud/issue.md","styling":{"logo":"https://b3tp9jag0.g.tfil.cloud/fws-full-logo.png","logo-mini":"https://b3tp9jag0.g.tfil.cloud/fws-small-logo.png"}}`))

	return 0
}
