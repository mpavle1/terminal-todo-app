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

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
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

		defer writer.Flush()

		indexForCompletion, err := strconv.Atoi(strings.Join(args, ""))
		if err != nil {
			log.Fatal(err)
		}

		lines, err := reader.ReadAll()

		if err != nil {
			log.Fatal(err)
		}

		var filteredLines [][]string

		for i := 0; i < len(lines); i++ {
			if i == 0 {
				filteredLines = append(filteredLines, lines[i])
				continue
			}
			index, err := strconv.Atoi(lines[i][0])
			if err != nil {
				log.Fatal(err)
			}
			if index == indexForCompletion {
				newLine := lines[i][:len(lines[i])-1]
				newLine = append(newLine, "true")
				filteredLines = append(filteredLines, newLine)
			} else {
				filteredLines = append(filteredLines, lines[i])
			}
		}

		file, err = os.Create("./data.csv")

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		writer = csv.NewWriter(file)

		err = writer.WriteAll(filteredLines)

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
