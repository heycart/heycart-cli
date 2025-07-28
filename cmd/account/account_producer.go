package account

import (
	"github.com/spf13/cobra"
)

var accountCompanyProducerCmd = &cobra.Command{
	Use:   "producer",
	Short: "Manage your HeyCart manufacturer",
}

func init() {
	accountRootCmd.AddCommand(accountCompanyProducerCmd)
}
