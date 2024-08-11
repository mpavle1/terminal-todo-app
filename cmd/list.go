/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open("./data.csv")

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		csvReader := csv.NewReader(f)

		data, err := csvReader.ReadAll()

		if err != nil {
			log.Fatal(err)
		}

		writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)

		defer writer.Flush()

		for i := range data {
			if i == 0 {
				fmt.Fprintln(writer, data[i])
				continue
			}
			var output []string
			for j := range data[i] {
				if j == 2 {
					todoCreatedTime, err := time.Parse(time.UnixDate, data[i][j])
					if err != nil {
						log.Fatal(err)
					}
					output = append(output, timediff.TimeDiff(todoCreatedTime))
					continue
				}
				output = append(output, data[i][j])
			}
			fmt.Fprintln(writer, output)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
