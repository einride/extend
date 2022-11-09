package main

import (
	"fmt"
	"os"

	saga "github.com/einride/saga/cmd/saga/gen"
	"github.com/einride/saga/cmd/saga/internal/auth"
	"github.com/spf13/cobra"
	"go.einride.tech/aip-cli/aipcli"
)

const (
	authFlagAnnotation = "saga_annotation_auth_flag"
	addressFlag        = "address"
	tokenFlag          = "token"
	clientIDFlag       = "client-id"
	clientSecretFlag   = "client-secret"
	authModuleName     = "auth"

	loginCmdName              = "login"
	logoutCmdName             = "logout"
	printIdentityTokenCmdName = "print-identity-token"
	exchangeSecretCmdName     = "exchange-secret"
)

func main() {
	moduleCmd := saga.NewModuleCommand(
		"saga",
		"Saga CLI",
		authModule(),
	)

	inputArgs := os.Args[1:]
	invokedCmd, _, err := moduleCmd.Find(inputArgs)
	if err != nil {
		exitOnError(err)
	}

	// Keep any PersistentPreRun functions set by aip-cli-go
	originalPreRun := invokedCmd.PersistentPreRun
	invokedCmd.PersistentPreRun = nil
	originalPreRunE := invokedCmd.PersistentPreRunE
	invokedCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		// no action for local, and unauthenticated commands
		if invokedCmd.Name() == exchangeSecretCmdName || (invokedCmd.Parent() != nil && invokedCmd.Parent().Name() == authModuleName) {
			return nil
		}
		if originalPreRunE != nil {
			if err := originalPreRunE(cmd, args); err != nil {
				return err
			}
		} else if originalPreRun != nil {
			originalPreRun(cmd, args)
		}

		// use the provided token if set
		if flag := invokedCmd.PersistentFlags().Lookup(tokenFlag); flag != nil && flag.Value.String() != "" {
			return nil
		}

		// authenticate with the same host as api
		host, err := extractHost(cmd)
		if err != nil {
			return err
		}

		accessToken, err := auth.Authenticate(invokedCmd.Context(), invokedCmd, host, aipcli.IsVerbose(invokedCmd))
		if err != nil {
			return err
		}
		err = invokedCmd.PersistentFlags().Set(tokenFlag, accessToken)
		if err != nil {
			return err
		}
		return nil
	}

	if err := moduleCmd.Execute(); err != nil {
		exitOnError(err)
	}
}

func exitOnError(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

// enforces that an api host must be set, either in address or host flag
func extractHost(cmd *cobra.Command) (string, error) {
	// check address flag
	if flagAddress, err := cmd.Flags().GetString(addressFlag); err == nil && flagAddress != "" {
		return flagAddress, nil
	}

	// check host flags
	for host, hostAddress := range aipcli.GetConfig(cmd).Hosts {
		if useHost, err := cmd.Flags().GetBool(host); err == nil && useHost {
			return hostAddress, nil
		}
	}

	// fail if not set
	return "", fmt.Errorf("missing address or host flag")
}

func authModule() *cobra.Command {
	config := saga.NewConfig()
	return aipcli.NewModuleCommand(
		authModuleName,
		"authentication commands for cli",
		config,
		loginCommand(),
		logoutCommand(),
		printIdentityTokenCommand(),
	)
}

func loginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   loginCmdName,
		Short: "Set user credentials in file saga/auth.json, used in authentication with Saga cli",
	}
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)

	cmd.Flags().String(clientIDFlag, "", "client ID to authenticate with")
	cmd.Flags().Lookup(clientIDFlag).Annotations = map[string][]string{
		authFlagAnnotation: {},
	}
	cmd.Flags().String(clientSecretFlag, "", "client secret to authenticate with")
	cmd.Flags().Lookup(clientSecretFlag).Annotations = map[string][]string{
		authFlagAnnotation: {},
	}

	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		var authFile auth.SagaAuthFile
		if !cmd.Flags().Changed(clientIDFlag) || !cmd.Flags().Changed(clientSecretFlag) {
			return fmt.Errorf("missing client-id and/or client-secret flags")
		}

		clientID, err := cmd.Flags().GetString(clientIDFlag)
		if err != nil {
			return err
		}
		clientSecret, err := cmd.Flags().GetString(clientSecretFlag)
		if err != nil {
			return err
		}

		authFile = auth.SagaAuthFile{
			ClientID:     clientID,
			ClientSecret: clientSecret,
		}
		if err := auth.StoreCredentials(cmd, &authFile); err != nil {
			return err
		}
		return nil
	}
	return cmd
}

func logoutCommand() *cobra.Command {
	authFilepath, err := auth.ConfigFilePath()
	if err != nil {
		exitOnError(err)
	}

	cmd := &cobra.Command{
		Use:   logoutCmdName,
		Short: fmt.Sprintf("Clear user credential file: %s", authFilepath),
	}
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)

	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		if err := auth.RemoveCredentials(cmd); err != nil {
			return err
		}
		return nil
	}
	return cmd
}

func printIdentityTokenCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   printIdentityTokenCmdName,
		Short: "Print access token for logged in user",
	}
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)

	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		host, err := extractHost(cmd)
		if err != nil {
			return err
		}

		token, err := auth.Authenticate(cmd.Context(), cmd, host, aipcli.IsVerbose(cmd))
		if err != nil {
			return err
		}
		cmd.Print(token)
		return nil
	}
	return cmd
}
