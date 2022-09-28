// Code generated by protoc-gen-go-cli. DO NOT EDIT.
// versions:
// 	protoc        (unknown)
package main

import (
	v1beta1 "github.com/einride/saga/cmd/saga/einride/extend/book/v1beta1"
	cobra "github.com/spf13/cobra"
	cli "go.einride.tech/protoc-gen-go-cli/cli"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "saga",
	}
	cmd.AddCommand(v1beta1.NewShipmentServiceCommand("shipment"))
	return cmd
}

func NewConfig() *cli.Config {
	return &cli.Config{
		Compiler: cli.CompilerConfig{Hosts: map[string]string{"prod": "api.saga.einride.tech"}, DefaultHost: "prod", Root: "saga", GoogleCloudIdentityTokens: false, Main: true},
	}
}
