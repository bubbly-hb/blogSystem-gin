package cmd

import (
	"os"

	"github.com/bubbly-hb/blogSystem-gin/config"
	"github.com/bubbly-hb/blogSystem-gin/db"
	"github.com/bubbly-hb/blogSystem-gin/model"
	"github.com/bubbly-hb/blogSystem-gin/router"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{}
)

func initConfig() {
	config.MustInit(os.Stdout, cfgFile) // 配置初始化
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config/dev.yaml", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().Bool("debug", true, "开启debug")
	viper.SetDefault("gin.mode", rootCmd.PersistentFlags().Lookup("debug"))
}

func Execute() error {
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		_, err := db.Mysql(
			viper.GetString("db.hostname"),
			viper.GetString("db.username"),
			viper.GetString("db.password"),
			viper.GetString("db.dbname"),
			viper.GetInt("db.port"),
		)
		if err != nil {
			return err
		}

		db.DB.AutoMigrate(&model.User{})

		defer db.DB.Close()

		r := router.SetupRouter()
		port := viper.GetString("server.port")
		return r.Run(port)
	}

	return rootCmd.Execute()

}
