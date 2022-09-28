package main

import (
	"context"
	"os"

	"go.einride.tech/sage/sg"
	"go.einride.tech/sage/sgtool"
	"go.einride.tech/sage/tools/sgapilinter"
	"go.einride.tech/sage/tools/sgbuf"
	"go.einride.tech/sage/tools/sgprotoc"
	"go.einride.tech/sage/tools/sgprotocgengogrpc"
)

type Proto sg.Namespace

func (Proto) Default(ctx context.Context) error {
	sg.Deps(ctx, Proto.BufFormat, Proto.BufLint, Proto.BufReadme)
	sg.Deps(ctx, Proto.APILinterLint)
	sg.Deps(ctx, Proto.BufGenerate, Proto.BufGenerateBook, Proto.BufGenerateCLI)
	return nil
}

func (Proto) BufLint(ctx context.Context) error {
	sg.Logger(ctx).Println("linting proto files...")
	cmd := sgbuf.Command(ctx, "lint")
	cmd.Dir = sg.FromGitRoot("proto")
	return cmd.Run()
}

func (Proto) APILinterLint(ctx context.Context) error {
	sg.Logger(ctx).Println("linting gRPC APIs...")
	return sgapilinter.Run(ctx)
}

func (Proto) BufFormat(ctx context.Context) error {
	sg.Logger(ctx).Println("linting proto files...")
	cmd := sgbuf.Command(ctx, "format", "--write")
	cmd.Dir = sg.FromGitRoot("proto")
	return cmd.Run()
}

func (Proto) BufPush(ctx context.Context) error {
	sg.Logger(ctx).Println("pushing Buf module..")
	cmd := sgbuf.Command(ctx, "push")
	cmd.Dir = sg.FromGitRoot("proto")
	return cmd.Run()
}

func (Proto) BufReadme(ctx context.Context) error {
	sg.Logger(ctx).Println("generating buf.md...")
	data, err := os.ReadFile(sg.FromGitRoot("docs", "apis.md"))
	if err != nil {
		return err
	}
	return os.WriteFile(sg.FromGitRoot("proto", "buf.md"), data, 0600)
}

func (Proto) ProtocGenGo(ctx context.Context) error {
	sg.Logger(ctx).Println("installing...")
	_, err := sgtool.GoInstallWithModfile(
		ctx,
		"google.golang.org/protobuf/cmd/protoc-gen-go",
		sg.FromGitRoot("cmd", "saga", "go.mod"),
	)
	return err
}

func (Proto) ProtocGenOpenAPIV2(ctx context.Context) error {
	sg.Logger(ctx).Println("installing...")
	_, err := sgtool.GoInstallWithModfile(
		ctx,
		"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2",
		sg.FromGitRoot("cmd", "saga", "go.mod"),
	)
	return err
}

func (Proto) ProtocGenGoAipCli(ctx context.Context) error {
	sg.Logger(ctx).Println("installing...")
	_, err := sgtool.GoInstallWithModfile(
		ctx,
		"go.einride.tech/aip-cli/cmd/protoc-gen-go-aip-cli",
		sg.FromGitRoot("cmd", "saga", "go.mod"),
	)
	return err
}

func (Proto) BufGenerate(ctx context.Context) error {
	sg.Deps(ctx, Proto.ProtocGenOpenAPIV2)
	sg.Logger(ctx).Println("generating proto stubs...")
	if err := os.RemoveAll(sg.FromGitRoot("openapiv2")); err != nil {
		return err
	}
	cmd := sgbuf.Command(
		ctx,
		"generate",
		"--output",
		sg.FromGitRoot(),
		"--template",
		"buf.gen.yaml",
		"--path",
		"einride/extend/book/v1beta1",
	)
	cmd.Dir = sg.FromGitRoot("proto")
	return cmd.Run()
}

func (Proto) BufGenerateBook(ctx context.Context) error {
	sg.Deps(ctx, Proto.ProtocGenOpenAPIV2)
	sg.Logger(ctx).Println("generating proto stubs...")
	cmd := sgbuf.Command(
		ctx,
		"generate",
		"--output",
		sg.FromGitRoot(),
		"--template",
		"buf.gen.book.yaml",
		"--path",
		"einride/extend/book/v1beta1",
	)
	cmd.Dir = sg.FromGitRoot("proto")
	return cmd.Run()
}

func (Proto) BufGenerateCLI(ctx context.Context) error {
	sg.Deps(
		ctx,
		sgprotoc.PrepareCommand,
		sgprotocgengogrpc.PrepareCommand,
		Proto.ProtocGenGo,
		Proto.ProtocGenGoAipCli,
	)
	sg.Logger(ctx).Println("generating protobuf stubs...")
	if err := os.RemoveAll(sg.FromGitRoot("cmd", "saga", "gen")); err != nil {
		return err
	}
	cmd := sgbuf.Command(
		ctx,
		"generate",
		"--output",
		sg.FromGitRoot(),
		"--template",
		"buf.gen.cli.yaml",
		"--path",
		"einride/extend/book/v1beta1",
	)
	cmd.Dir = sg.FromGitRoot("proto")
	if err := cmd.Run(); err != nil {
		return err
	}
	cmd = sg.Command(ctx, "go", "mod", "tidy")
	cmd.Dir = sg.FromGitRoot("cmd", "saga")
	return cmd.Run()
}
