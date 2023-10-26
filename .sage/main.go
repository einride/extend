package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"go.einride.tech/sage/sg"
	"go.einride.tech/sage/tools/sgconvco"
	"go.einride.tech/sage/tools/sggit"
	"go.einride.tech/sage/tools/sggolangcilint"
	"go.einride.tech/sage/tools/sgmdformat"
	"go.einride.tech/sage/tools/sgyamlfmt"
)

func main() {
	sg.GenerateMakefiles(
		sg.Makefile{
			Path:          sg.FromGitRoot("Makefile"),
			DefaultTarget: Default,
		},

		sg.Makefile{
			Path:          sg.FromGitRoot("proto", "Makefile"),
			Namespace:     Proto{},
			DefaultTarget: Proto.Default,
		},
	)
}

func Default(ctx context.Context) error {
	sg.Deps(ctx, ConvcoCheck, FormatMarkdown, FormatYaml)
	sg.Deps(ctx, Proto.Default)
	sg.Deps(ctx, GoModTidy)
	sg.Deps(ctx, GoLint)
	sg.Deps(ctx, GitVerifyNoDiff)
	return nil
}

func GoModTidy(ctx context.Context) error {
	sg.Logger(ctx).Println("tidying Go module files...")
	command := sg.Command(ctx, "go", "mod", "tidy", "-v")
	command.Dir = sg.FromGitRoot()
	return command.Run()
}

func ConvcoCheck(ctx context.Context) error {
	sg.Logger(ctx).Println("checking git commits...")
	return sgconvco.Command(ctx, "check", "origin/master..HEAD").Run()
}

func FormatMarkdown(ctx context.Context) error {
	sg.Logger(ctx).Println("formatting Markdown files...")
	return sgmdformat.Command(ctx).Run()
}

func FormatYaml(ctx context.Context) error {
	sg.Logger(ctx).Println("formatting YAML files...")
	return sgyamlfmt.Run(ctx)
}

func GitVerifyNoDiff(ctx context.Context) error {
	sg.Logger(ctx).Println("verifying that git has no diff...")
	return sggit.VerifyNoDiff(ctx)
}

func GoLint(ctx context.Context) error {
	sg.Logger(ctx).Println("linting Go files...")
	return sggolangcilint.CommandInDirectory(
		ctx,
		sg.FromGitRoot(),
	).Run()
}

func ServePDF(ctx context.Context) error {
	const addr = ":8080"
	httpMux := http.NewServeMux()
	httpMux.Handle("/openapiv2/", http.StripPrefix("/openapiv2", http.FileServer(http.Dir("./openapiv2"))))
	httpMux.Handle("/", http.FileServer(http.Dir("./pdf")))
	sg.Logger(ctx).Printf("listening on %s", addr)
	return (&http.Server{
		Addr:    addr,
		Handler: httpMux,
	}).ListenAndServe()
}

func GenerateApiDoc(ctx context.Context) error {
	sg.SerialDeps(ctx, installRapiPdfCli)
	filePath, err := createTempFile()
	if err != nil {
		return err
	}
	sg.SerialDeps(ctx, sg.Fn(replaceDatePlaceholderWithActual, filePath))
	sg.Deps(
		ctx,
		sg.Fn(generatePdfFromSwagger, "auth", filePath),
		sg.Fn(generatePdfFromSwagger, "book", filePath),
	)
	return nil
}

func installRapiPdfCli(ctx context.Context) error {
	if err := sg.Command(ctx, "yarn", "global", "add", "node-gyp").Run(); err != nil {
		return err
	}
	cmd := sg.Command(ctx, "yarn", "global", "add", "@kingjan1999/rapipdf-cli")
	return cmd.Run()
}

func createTempFile() (string, error) {
	// Create a new temporary file that will contain the modified rapipdf config.
	temp, err := os.CreateTemp("", "rapipdf.copy.json")
	if err != nil {
		return "", err
	}
	filePath := temp.Name()
	if err := temp.Close(); err != nil {
		return "", err
	}
	return filePath, nil
}

func replaceDatePlaceholderWithActual(ctx context.Context, configPath string) error {
	sg.Logger(ctx).Println("replacing placeholder with today's date...")
	orginalConfigBytes, err := os.ReadFile(sg.FromGitRoot("pdf/rapipdf.json"))
	if err != nil {
		return err
	}
	originalConfig := string(orginalConfigBytes)
	date := time.Now().Format("2006-01-02")
	updatedConfig := strings.Replace(originalConfig, "<DATE_PLACEHOLDER>", date, 1)
	f, err := os.OpenFile(configPath, os.O_WRONLY, 0o755)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	_, err = f.Write([]byte(updatedConfig))
	return err
}

func generatePdfFromSwagger(ctx context.Context, name, configPath string) error {
	filename := fmt.Sprintf("api_%s_latest.pdf", name)
	api := fmt.Sprintf("openapiv2/%s.swagger.yaml", name)
	return generatePdf(ctx, configPath, filename, api)
}

func generatePdf(ctx context.Context, configPath, outputFilename, openApiFileName string) error {
	cmd := sg.Command(
		ctx,
		"rapipdf",
		"--configFile="+configPath,
		"--outputFile="+outputFilename,
		openApiFileName,
	)
	return cmd.Run()
}
