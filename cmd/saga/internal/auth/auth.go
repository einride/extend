package auth

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/adrg/xdg"
	authv1beta1 "github.com/einride/saga/cmd/saga/gen/einride/saga/extend/auth/v1beta1"
	"github.com/spf13/cobra"
	"go.einride.tech/aip-cli/aipcli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	authConfigFile = "saga/auth.json"
)

func Authenticate(ctx context.Context, cmd *cobra.Command, host string, verboseLogging bool) (string, error) {
	authFile, err := readCredentials()
	if err != nil {
		return "", err
	}
	clientID := authFile.ClientID
	clientSecret := authFile.ClientSecret

	if verboseLogging {
		cmd.Printf(">> using client-id: %s, and client-secret with length: %d \n", clientID, len(clientSecret))
	}

	authClient, cleanup, err := authServiceClient(ctx, host)
	if err != nil {
		return "", err
	}
	defer cleanup()

	request := authv1beta1.ExchangeSecretRequest{
		ClientId:     clientID,
		ClientSecret: clientSecret,
	}

	response, err := authClient.ExchangeSecret(ctx, &request)
	if err != nil {
		return "", err
	}
	return response.AccessToken, nil
}

type SagaAuthFile struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

func readCredentials() (*SagaAuthFile, error) {
	authFilepath, err := ConfigFilePath()
	if err != nil {
		return nil, err
	}
	if _, err = os.Stat(authFilepath); err != nil {
		return nil, fmt.Errorf("no logged in user, use 'saga auth login " +
			"--client-id <CLIENT_ID> --client-secret <CLIENT_SECRET>'")
	}
	data, err := os.ReadFile(authFilepath)
	if err != nil {
		return nil, err
	}
	var authFile SagaAuthFile
	if err := json.Unmarshal(data, &authFile); err != nil {
		return nil, err
	}
	return &authFile, nil
}

func StoreCredentials(cmd *cobra.Command, authFile *SagaAuthFile) error {
	tokenData, err := json.MarshalIndent(&authFile, "", "  ")
	if err != nil {
		return err
	}
	authFilepath, err := ConfigFilePath()
	if err != nil {
		return err
	}
	if err := os.WriteFile(authFilepath, tokenData, 0o600); err != nil {
		return err
	}
	if aipcli.IsVerbose(cmd) {
		cmd.Printf("successfully stored credentials in file: %s\n", authFilepath)
	}
	return nil
}

func RemoveCredentials(cmd *cobra.Command) error {
	authFilepath, err := ConfigFilePath()
	if err != nil {
		return err
	}
	if err := os.Remove(authFilepath); err != nil {
		return err
	}
	if aipcli.IsVerbose(cmd) {
		cmd.Printf("successfully removed file: %s\n", authFilepath)
	}
	return nil
}

func authServiceClient(ctx context.Context, target string) (authv1beta1.AuthenticationServiceClient, func(), error) {
	systemCertPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, nil, fmt.Errorf("dial %s: %w", target, err)

	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(systemCertPool, "")),
		grpc.WithBlock(),
	}
	conn, err := grpc.DialContext(ctx, withDefaultPort(target, 443), opts...)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := conn.Close(); err != nil {
			fmt.Printf("close service connection: %s", err.Error())
		}
	}
	return authv1beta1.NewAuthenticationServiceClient(conn), cleanup, nil
}

func withDefaultPort(target string, port int) string {
	parts := strings.Split(target, ":")
	if len(parts) == 1 {
		return target + ":" + strconv.Itoa(port)
	}
	return target
}

func ConfigFilePath() (string, error) {
	return xdg.ConfigFile(authConfigFile)
}
