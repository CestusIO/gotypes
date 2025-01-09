//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package types

import "github.com/google/wire"

// IDProviderSet is a provider set for an UUIDProvider
var IDProviderSet = wire.NewSet(
	ProvideIDProvider,
	wire.Bind(new(IDProvider), new(*defaultIDProvider)),
)

// TimeProviderSet is a provider set for an TimeProvider
var TimeProviderSet = wire.NewSet(
	ProvideTimeProvider,
	wire.Bind(new(TimeProvider), new(*defaultTimeProvider)),
)
