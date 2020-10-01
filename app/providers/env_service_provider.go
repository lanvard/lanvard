package providers

import (
	"github.com/lanvard/contract/inter"
	"lanvard/config"
)

type EnvServiceProvider struct{}

// Register any container services.
func (a EnvServiceProvider) Register(container inter.Container) inter.Container {
	container.Instance("env", config.App.Env)

	return container
}