package sgprotocgenjsonschema

import (
	"context"
	"os/exec"

	"go.einride.tech/sage/sg"
	"go.einride.tech/sage/sgtool"
)

const (
	version = "1.3.10"
	name    = "protoc-gen-jsonschema"
)

func Command(ctx context.Context, args ...string) *exec.Cmd {
	sg.Deps(ctx, PrepareCommand)
	return sg.Command(ctx, sg.FromBinDir(name), args...)
}

func PrepareCommand(ctx context.Context) error {
	_, err := sgtool.GoInstall(ctx, "github.com/chrusty/protoc-gen-jsonschema/cmd/protoc-gen-jsonschema", "latest")
	if err != nil {
		return err
	}
	return nil
}
