package cmd

import (
	"github.com/lang-dev/exchange/config"
	"github.com/lang-dev/exchange/internal/entity"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	mCommand = &cobra.Command{
		Use:   "migrate",
		Short: "Initializes the database",
		Long:  "Initializes the database",
		RunE:  mExecute,
	}
)

func init() {
	rootCmd.AddCommand(mCommand)
}

func mExecute(cmd *cobra.Command, args []string) error {
	db, err := gorm.Open(sqlite.Open(config.DatabaseName), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{})
	return nil
}
