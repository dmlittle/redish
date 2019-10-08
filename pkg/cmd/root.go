package cmd

import (
	"github.com/spf13/cobra"
)

// Flags is a struct that holds the values of any potential flag passed into
// `redish`.
type Flags struct {
	Auth     string
	Database int
	Hostname string
	Port     int
	URI      string
	TLS      bool
}

// NewRedishCommand creates the `redish` command.
func NewRedishCommand() *cobra.Command {
	f := &Flags{}

	cmd := &cobra.Command{
		Use:   "redish",
		Short: "a Redis CLI with built-in TLS support.",
		Run:   func(_ *cobra.Command, _ []string) {},
	}

	cmd.Flags().StringVarP(&f.Auth, "auth", "a", "", "password to use when connecting to the server.")
	cmd.Flags().IntVarP(&f.Database, "database", "n", 0, "database number.")
	cmd.Flags().StringVarP(&f.Hostname, "host", "H", "127.0.0.1", "server hostname.")
	cmd.Flags().IntVarP(&f.Port, "port", "p", 6379, "server port.")
	cmd.Flags().StringVarP(&f.URI, "uri", "u", "", "server URI.")
	cmd.Flags().BoolVar(&f.TLS, "tls", false, "use TLS connection.")

	return cmd
}
