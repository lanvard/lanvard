package decorator

import (
	"github.com/lanvard/contract/inter"
	"lanvard/app/providers"
)

type BootProviders struct{}

// Providers are located in config/providers/providers.go
func (r BootProviders) Bootstrap(container inter.Container) inter.Container {
	for _, bootstrapper := range providers.Providers.BootProviders {
		container = bootstrapper.Boot(container)
	}

	return container
}
