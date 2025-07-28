package project

import (
	"fmt"
	"os"

	adminSdk "github.com/heycart/go-heycart-admin-api-sdk"
	"github.com/spf13/cobra"

	"github.com/heycart/heycart-cli/logging"
	"github.com/heycart/heycart-cli/shop"
)

var projectClearCacheCmd = &cobra.Command{
	Use:   "clear-cache",
	Short: "Clears the Shop cache",
	RunE: func(cmd *cobra.Command, _ []string) error {
		var cfg *shop.Config
		var err error

		if cfg, err = shop.ReadConfig(projectConfigPath, false); err != nil {
			return err
		}

		if cfg.AdminApi == nil {
			logging.FromContext(cmd.Context()).Infof("Clearing cache localy")

			projectRoot, err := findClosestHeyCartProject()
			if err != nil {
				return err
			}

			return os.RemoveAll(fmt.Sprintf("%s/var/cache", projectRoot))
		}

		logging.FromContext(cmd.Context()).Infof("Clearing cache using admin-api")

		client, err := shop.NewShopClient(cmd.Context(), cfg)
		if err != nil {
			return err
		}

		_, err = client.CacheManager.Clear(adminSdk.NewApiContext(cmd.Context()))

		return err
	},
}

func init() {
	projectRootCmd.AddCommand(projectClearCacheCmd)
}
