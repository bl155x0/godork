package cmd

import (
	"bufio"
	"godork/dork"
	"os"

	"github.com/spf13/cobra"
)

const (
	GODORK_DORKFILE_ENV = "GODORK_DORKFILE"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "godork",
	Short: "A simple tool for creating google dork search links for a bunch of given domains.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		dorkfile, isSet := os.LookupEnv(GODORK_DORKFILE_ENV)
		if isSet == false || dorkfile == "" {
			cmd.MarkFlagRequired("dorkfile")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		dorkfile := getDorkFile(cmd)
		url, err := cmd.Flags().GetString("url")
		cobra.CheckErr(err)

		dorks, err := createDorks(dorkfile, url)
		cobra.CheckErr(err)

		processor, err := getProcessor(cmd)
		cobra.CheckErr(err)

		err = processDorks(processor, dorks)
		cobra.CheckErr(err)
	},
}

// createDorks from the given file using the given URL
func createDorks(dorkfile string, url string) ([]string, error) {
	var dorks []string
	if url != "" {
		//append a single URL from the url flag
		err := dork.AppendDorkForUrl(url, dorkfile, &dorks)
		if err != nil {
			return nil, err
		}
	} else {
		//append a bunch of URLs read from STDIN
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			url := scanner.Text()
			err := dork.AppendDorkForUrl(url, dorkfile, &dorks)
			if err != nil {
				return nil, err
			}
		}
	}
	return dorks, nil
}

// getProcessor reports the default processor which should be used to process the dorks
func getProcessor(cmd *cobra.Command) (*dork.DorkPrinter, error) {
	urlencode, err := cmd.Flags().GetBool("urlencode")
	if err != nil {
		return nil, err
	}
	return &dork.DorkPrinter{
		UrlEncode: urlencode,
	}, nil
}

// processDorks will process the dorks with the given processor
func processDorks(p dork.DorkProcessor, dorks []string) error {
	return p.Process(dorks)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// getDorkFile reports the dork file to use. Either from the ENV variable, or the command line flag
func getDorkFile(cmd *cobra.Command) string {
	dorkfile := os.Getenv(GODORK_DORKFILE_ENV)
	if dorkfile != "" {
		return dorkfile
	}

	dorkfile, err := cmd.Flags().GetString("dorkfile")
	cobra.CheckErr(err)
	return dorkfile
}

func init() {
	rootCmd.Flags().BoolP("urlencode", "e", false, "URL-Encode the output URLs")
	rootCmd.Flags().StringP("dorkfile", "d", "", "Dorks file to use (or set env GODORK_DORKFILE)")
	rootCmd.Flags().StringP("url", "u", "", "a single URL to ceate the dorks for")
}
