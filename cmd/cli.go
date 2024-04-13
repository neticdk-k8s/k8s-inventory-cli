package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/neticdk-k8s/ic/internal/apiclient"
	"github.com/neticdk-k8s/ic/internal/logger"
	"github.com/neticdk-k8s/ic/internal/oidc"
	"github.com/neticdk-k8s/ic/internal/reader"
	"github.com/neticdk-k8s/ic/internal/tokencache"
	"github.com/neticdk-k8s/ic/internal/ui"
	"github.com/neticdk-k8s/ic/internal/usecases/authentication"
	"github.com/neticdk-k8s/ic/internal/usecases/authentication/authcode"
	"golang.org/x/term"
)

type OIDCConfig struct {
	IssuerURL                   string
	ClientID                    string
	GrantType                   string
	RedirectURLHostname         string
	RedirectURIAuthCodeKeyboard string
	AuthBindAddr                string
	TokenCacheDir               string
}

type ExecutionContext struct {
	Stderr, Stdout io.Writer

	// Logger is the global logger object to print logs.
	Logger logger.Logger

	// Spinner is the global spinner object used to show progress across the cli
	Spinner *ui.Spinner

	// LogLevel is the log level used for the logger
	LogLevel string

	// APIServer is the inventory api server endpoint
	APIServer string

	// Interactive can be used to force interactive/non-interactive use
	Interactive string

	// IsTerminal indicates whether the current session is a terminal or not
	IsTerminal bool

	// OutputFormat is the format used for outputting data
	OutputFormat string

	OIDC          OIDCConfig
	OIDCProvider  oidc.Provider
	Authenticator authentication.Authenticator
	TokenCache    tokencache.Cache

	APIClient *apiclient.ClientWithResponses
}

func NewExecutionContext() *ExecutionContext {
	ec := &ExecutionContext{
		Stderr:       os.Stderr,
		Stdout:       os.Stdout,
		OutputFormat: "text",
		OIDC:         OIDCConfig{},
	}
	return ec
}

func (ec *ExecutionContext) SetupAPIClient(token string) error {
	var err error
	provider := apiclient.NewAuthProvider(token)
	ec.APIClient, err = apiclient.NewClientWithResponses(
		ec.APIServer,
		apiclient.WithRequestEditorFn(provider.Intercept))
	if err != nil {
		return err
	}
	return nil
}

// Prepare sets up context that does not depend on flags
// It should be called before rootCmd.Execute
func (ec *ExecutionContext) Prepare() error {
	ec.IsTerminal = term.IsTerminal(int(os.Stdout.Fd()))

	ec.SetupLogger("info")

	ec.setupSpinner()

	return nil
}

// Setup sets up context that depends on the flags being sets
// It should be called from rootCmd.PersistentPreRunE
func (ec *ExecutionContext) Setup() error {
	ec.SetupLogger(ec.LogLevel)

	authn := authentication.NewAuthentication(
		ec.Logger,
		nil,
		&authcode.Browser{Logger: ec.Logger},
		&authcode.Keyboard{Reader: reader.NewReader(), Logger: ec.Logger})
	ec.Authenticator = authentication.NewAuthenticator(ec.Logger, authn)

	ec.OIDCProvider = oidc.Provider{
		IssuerURL:   ec.OIDC.IssuerURL,
		ClientID:    ec.OIDC.ClientID,
		ExtraScopes: []string{"profile", "email", "roles", "offline_access"},
	}

	var err error
	if ec.TokenCache, err = tokencache.NewFSCache(ec.OIDC.TokenCacheDir); err != nil {
		return fmt.Errorf("creating token cache: %w", err)
	}

	return nil
}

func (ec *ExecutionContext) setupSpinner() {
	if ec.Spinner == nil {
		ec.Spinner = ui.NewSpinner(ec.Stderr, ec.Logger)
	}
}

func (ec *ExecutionContext) Spin(message string) {
	if !ec.Spinner.Running() {
		ec.Spinner.Run()
	}
	ec.Spinner.Text(message)
}

func (ec *ExecutionContext) SetupLogger(logLevel string) {
	if ec.Logger == nil {
		ec.Logger = logger.New(ec.Stderr, logLevel)
	} else {
		if err := ec.Logger.SetLevel(logLevel); err != nil {
			ec.Logger.Error("Failed to set loglevel", "level", logLevel)
		}
	}
	if logLevel != ec.LogLevel {
		ec.LogLevel = logLevel
	}
	ec.Logger.SetInteractive(ec.Interactive, ec.IsTerminal)
}
