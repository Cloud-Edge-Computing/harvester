package api

import (
	"context"

	"github.com/rancher/steve/pkg/server"

	"github.com/harvester/harvester/pkg/api/datavolume"
	"github.com/harvester/harvester/pkg/api/image"
	"github.com/harvester/harvester/pkg/api/keypair"
	"github.com/harvester/harvester/pkg/api/network"
	"github.com/harvester/harvester/pkg/api/node"
	"github.com/harvester/harvester/pkg/api/restore"
	"github.com/harvester/harvester/pkg/api/setting"
	"github.com/harvester/harvester/pkg/api/upgrade"
	"github.com/harvester/harvester/pkg/api/user"
	"github.com/harvester/harvester/pkg/api/vm"
	"github.com/harvester/harvester/pkg/api/vmtemplate"
	"github.com/harvester/harvester/pkg/config"
)

type registerSchema func(scaled *config.Scaled, server *server.Server, options config.Options) error

func registerSchemas(scaled *config.Scaled, server *server.Server, options config.Options, registers ...registerSchema) error {
	for _, register := range registers {
		if err := register(scaled, server, options); err != nil {
			return err
		}
	}
	return nil
}

func Setup(ctx context.Context, server *server.Server, controllers *server.Controllers, options config.Options) error {
	scaled := config.ScaledWithContext(ctx)
	return registerSchemas(scaled, server, options,
		image.RegisterSchema,
		keypair.RegisterSchema,
		vmtemplate.RegisterSchema,
		vm.RegisterSchema,
		setting.RegisterSchema,
		upgrade.RegisterSchema,
		user.RegisterSchema,
		network.RegisterSchema,
		node.RegisterSchema,
		datavolume.RegisterSchema,
		restore.RegisterSchema)
}
