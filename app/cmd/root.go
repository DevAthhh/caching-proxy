package cmd

import (
	"fmt"
	"log"

	"github.com/DevAthhh/caching-proxy/app/cmd/caching"
	"github.com/DevAthhh/caching-proxy/app/cmd/server"
	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	var port string
	var origin string

	cmd := &cobra.Command{
		Use: "caching-proxy",
		Run: func(cmd *cobra.Command, args []string) {
			cls, err := cmd.Flags().GetBool("cls-cache")
			if err != nil {
				log.Fatalf("error with flag 'clear-cache': %v", err)
			}
			if cls {
				if err := caching.ClearAllCache(); err != nil {
					log.Fatalf("error with clearing cache: %v", err)
				}
				fmt.Println("the cache has been cleared")
			} else {
				server.Run(port, origin)
			}

		},
	}

	cmd.Flags().StringVarP(&port, "port", "p", "5050", "port flag")
	cmd.Flags().StringVarP(&origin, "origin", "o", "", "origin flag")
	cmd.Flags().Bool("cls-cache", false, "clear-cache flag")

	return cmd
}
