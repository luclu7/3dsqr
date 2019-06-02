package cmd

import (
	"fmt"
	"github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/spf13/cobra"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Errorf("requires a color argument")
		}
		match, _ := regexp.MatchString("[a-zA-Z0-9]+\\.cia$", args[0])

		if match == false {
			fmt.Println("Please provide a CIA file.")
			os.Exit(1)
		}

		ip := GetOutboundIP()
		ipString := ip.String()
		content := "http://" + ipString + ":8000/" + args[0]
		obj := qrcodeTerminal.New()
		obj.Get([]byte(content)).Print()
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		http.Handle("/", http.FileServer(http.Dir(dir)))
		log.Printf("Serving your CIA at: " + content)
		log.Fatal(http.ListenAndServe(":8000", nil))

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
