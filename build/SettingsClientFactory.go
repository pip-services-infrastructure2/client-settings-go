package build

import (
	clients1 "github.com/pip-services-infrastructure2/client-settings-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type SettingsClientFactory struct {
	*cbuild.Factory
}

func NewSettingsClientFactory() *SettingsClientFactory {
	c := &SettingsClientFactory{
		Factory: cbuild.NewFactory(),
	}

	cmdHttpClientDescriptor := cref.NewDescriptor("service-accounts", "client", "commandable-http", "*", "1.0")
	cmdGrpcClientDescriptor := cref.NewDescriptor("service-accounts", "client", "commandable-gtpc", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("service-accounts", "client", "grpc", "*", "1.0")

	c.RegisterType(cmdHttpClientDescriptor, clients1.NewSettingsCommandableHttpClientV1)
	c.RegisterType(cmdGrpcClientDescriptor, clients1.NewSettingsCommandableGrpcClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewSettingsGrpcClientV1)

	return c
}
