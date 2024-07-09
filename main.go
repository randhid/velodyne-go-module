package main

import (
	"context"

	"go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/utils"

	"velodyne/velodyne"
)

func main() {
	utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("velodyne-go-module"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) (err error) {
	// instantiates the module itself
	mod, err := module.NewModuleFromArgs(ctx, logger)
	if err != nil {
		return err
	}

	// Camerass
	if err = mod.AddModelFromRegistry(ctx, camera.API, velodyne.Model); err != nil {
		return err
	}

	// Each module runs as its own process
	err = mod.Start(ctx)
	logger.Warn("starting velodyne-go-module")
	defer mod.Close(ctx)
	if err != nil {
		return err
	}
	<-ctx.Done()
	return nil
}
