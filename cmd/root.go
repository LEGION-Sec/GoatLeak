package cmd

import (
	"fmt"
	"os"

	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

)

const banner = `

  ▄████  ▒█████   ▄▄▄     ▄▄▄█████▓ ██▓    ▓█████ ▄▄▄       ██ ▄█▀
 ██▒ ▀█▒▒██▒  ██▒▒████▄   ▓  ██▒ ▓▒▓██▒    ▓█   ▀▒████▄     ██▄█▒ 
▒██░▄▄▄░▒██░  ██▒▒██  ▀█▄ ▒ ▓██░ ▒░▒██░    ▒███  ▒██  ▀█▄  ▓███▄░ 
░▓█  ██▓▒██   ██░░██▄▄▄▄██░ ▓██▓ ░ ▒██░    ▒▓█  ▄░██▄▄▄▄██ ▓██ █▄ 
░▒▓███▀▒░ ████▓▒░ ▓█   ▓██▒ ▒██▒ ░ ░██████▒░▒████▒▓█   ▓██▒▒██▒ █▄
 ░▒   ▒ ░ ▒░▒░▒░  ▒▒   ▓▒█░ ▒ ░░   ░ ▒░▓  ░░░ ▒░ ░▒▒   ▓▒█░▒ ▒▒ ▓▒
  ░   ░   ░ ▒ ▒░   ▒   ▒▒ ░   ░    ░ ░ ▒  ░ ░ ░  ░ ▒   ▒▒ ░░ ░▒ ▒░
░ ░   ░ ░ ░ ░ ▒    ░   ▒    ░        ░ ░      ░    ░   ▒   ░ ░░ ░ 
      ░     ░ ░        ░  ░            ░  ░   ░  ░     ░  ░░  ░   
                                                                  
		⠀⠀⠀⠀⠀⠀⠀⢀⣠⡤⣤⡀⠀⠀⠀⠀⠀⠀⢀⡠⡤⣄⣀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⢀⡴⠊⢉⣑⣽⣕⡈⢢⠀⠀⠀⠀⡰⠉⣢⣿⣃⡉⠉⠢⡀⠀⠀⠀⠀
		⠀⠀⠒⠶⠥⠒⠉⠁⠀⠀⢱⠒⠈⡆⠀⠀⢠⠃⠒⡜⠀⠀⠀⠉⠒⠨⠵⠖⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡧⠬⢽⠀⠀⡼⠭⢬⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢛⣙⠚⡆⠀⡗⣊⣹⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣸⡒⠫⢇⡸⠌⠒⣼⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⢀⠤⠤⢤⣒⠉⠁⡏⠀⠙⡦⠤⠤⢴⠞⠁⢸⠌⠉⢒⣤⠤⠤⡄⠀⠀⠀
		⠀⠀⠀⠘⢄⡀⠣⠤⣙⡿⣿⣕⡄⢸⠀⠀⡜⢀⡮⣽⠿⣛⠡⠜⢀⡠⠏⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠉⠁⠀⠀⠈⡛⢿⢸⠀⠀⡇⡼⠛⡇⠀⠀⠈⠉⠁⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢣⢸⢸⠀⠀⢃⡇⡰⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣇⡇⠆⠀⡜⣠⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢹⡔⠧⠴⠃⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣗⠤⠤⣺⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣏⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣻⣸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀


	    Creator: LEGION-Sec (shusig2603@gmail.com)

		Don't Let Secrets Escape the Herd!





`

const configDescription = `config file path (required)
Only the --config flag is supported for configuration`

var rootCmd = &cobra.Command{
	Use:   "goatleak",
	Short: "GoatLeak scans code, past or present, for secrets",
}

func init() {
	cobra.OnInitialize(initLog)
	rootCmd.PersistentFlags().StringP("config", "c", "", configDescription)
	rootCmd.PersistentFlags().Int("exit-code", 1, "exit code when leaks have been encountered")
	rootCmd.PersistentFlags().StringP("source", "s", ".", "path to source")
	rootCmd.PersistentFlags().StringP("report-path", "r", "", "report file")
	rootCmd.PersistentFlags().StringP("report-format", "f", "json", "output format (json, csv, junit, sarif)")
	rootCmd.PersistentFlags().StringP("baseline-path", "b", "", "path to baseline with issues that can be ignored")
	rootCmd.PersistentFlags().StringP("log-level", "l", "info", "log level (trace, debug, info, warn, error, fatal)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "show verbose output from scan")
	rootCmd.PersistentFlags().BoolP("no-color", "", false, "turn off color for verbose output")
	rootCmd.PersistentFlags().Int("max-target-megabytes", 0, "files larger than this will be skipped")
	rootCmd.PersistentFlags().Bool("redact", false, "redact secrets from logs and stdout")
	rootCmd.PersistentFlags().Bool("no-banner", false, "suppress banner")
	rootCmd.PersistentFlags().String("exceptions", "", "path to exceptions JSON file") 
	rootCmd.PersistentFlags().StringSlice("quality-gate", []string{}, "quality gate: severity=count (e.g., critical=0, high=5)") 
	err := viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	if err != nil {
		log.Fatal().Msgf("err binding config %s", err.Error())
	}
}

func initLog() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	ll, err := rootCmd.Flags().GetString("log-level")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	switch strings.ToLower(ll) {
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "err", "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

func initConfig() { 
	hideBanner, err := rootCmd.Flags().GetBool("no-banner")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	if !hideBanner {
		_, _ = fmt.Fprint(os.Stderr, banner)
	}
	
	cfgPath, err := rootCmd.Flags().GetString("config")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	
	// FORCE CONFIG REQUIREMENT - ONLY --config FLAG SUPPORTED
	if cfgPath == "" {
		log.Fatal().Msg("Error: --config flag is required. No default configuration is provided.")
	}
	
	// Check if config file exists
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatal().Msgf("Config file not found: %s", cfgPath)
	}
	
	// Load ONLY from provided config file
	viper.SetConfigFile(cfgPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("Failed to load config file")
	}
	
	log.Debug().Msgf("Using config: %s", cfgPath)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		if strings.Contains(err.Error(), "unknown flag") {
			// exit code 126: Command invoked cannot execute
			os.Exit(126)
		}
		log.Fatal().Msg(err.Error())
	}
}
