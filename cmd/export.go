package cmd

import (
	"fmt"
	"github.com/nyudlts/go-aspace/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var resourceId int
var repositoryId int
var location string
var client lib.ASClient

func init() {
	log.Println("init")
	client = lib.Client
	rootCmd.AddCommand(exportCmd)
	exportCmd.PersistentFlags().IntVar(&repositoryId, "repositoryId", 0, "Id of the repository")
	viper.BindPFlag("repositoryId", exportCmd.PersistentFlags().Lookup("repositoryId"))
	exportCmd.PersistentFlags().IntVar(&resourceId, "resourceId", 0, "Id of the resource")
	viper.BindPFlag("resourceId", exportCmd.PersistentFlags().Lookup("resourceId"))
	exportCmd.PersistentFlags().StringVar(&location, "location", "/tmp", "Location to write EAD File")
	viper.BindPFlag("location", exportCmd.PersistentFlags().Lookup("location"))

}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export a resource from archivesspace",
	Long:  `export a resource from archivesspace`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("** exporting")
		err := exportEAD(); if err != nil {
			panic(err)
		}
		fmt.Println("Export complete")
	},
}

func exportEAD() error {
	if _, err := os.Stat(viper.GetString("location")); os.IsNotExist(err) {
		return fmt.Errorf("%v does not exist, exiting", err.Error())
	}
	err := client.SerializeEAD(viper.GetInt("repositoryId"), viper.GetInt("resourceId"), viper.GetString("location")); if err != nil {
		return err
	}

	return nil
}

