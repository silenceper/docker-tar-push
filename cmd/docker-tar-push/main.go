package main

import (
	"github.com/silenceper/docker-tar-push/pkg/push"
	"github.com/silenceper/log"
	"github.com/spf13/cobra"
)

var (
	registryURL   string
	username      string
	password      string
	repo          string
	skipSSLVerify bool
	logLevel      int

	rootCmd = &cobra.Command{
		Use:   "docker-tar-push",
		Short: "push your docker tar archive image without docker",
		Long:  `push your docker tar archive image without docker.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			log.SetLogLevel(log.Level(logLevel))

			imagePush := push.NewImagePush(args[0], registryURL, username, password, repo, skipSSLVerify)
			imagePush.Push()
		},
	}
)

func main() {
	rootCmd.Flags().StringVar(&registryURL, "registry", "", "registry url")
	rootCmd.Flags().StringVar(&username, "username", "", "registry auth username")
	rootCmd.Flags().StringVar(&password, "password", "", "registry auth password")
	rootCmd.Flags().StringVar(&repo, "repo", "", "docker image repository")
	rootCmd.Flags().BoolVar(&skipSSLVerify, "skip-ssl-verify", false, "skip ssl verify")
	rootCmd.Flags().IntVar(&logLevel, "log-level", log.LevelInfo, "log-level, 0:Fatal,1:Error,2:Warn,3:Info,4:Debug")

	rootCmd.MarkFlagRequired("registry")

	rootCmd.Execute()
}
