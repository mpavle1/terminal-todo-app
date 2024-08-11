/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.OpenFile("./data.csv", os.O_APPEND|os.O_RDWR, 0644)

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		writer := csv.NewWriter(file)
		reader := csv.NewReader(file)
		stdOutWriter := csv.NewWriter(os.Stdout)

		defer writer.Flush()
		defer stdOutWriter.Flush()

		lines, err := reader.ReadAll()

		if err != nil {
			log.Fatal(err)
		}

		lastIndex := 0

		if len(lines) > 1 {
			lastIndex, err = strconv.Atoi(lines[len(lines)-1][0])
			if err != nil {
				log.Fatal(err)
			}
		}

		index := lastIndex + 1

		now := time.Now()

		now.Format(time.UnixDate)

		writer.Write([]string{
			strconv.Itoa(index), strings.Join(args, " "), now.String(), "false",
		})

		stdOutWriter.Write([]string{
			strconv.Itoa(index), strings.Join(args, " "), now.String(), "false",
		})
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
