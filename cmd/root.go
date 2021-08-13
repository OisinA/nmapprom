package cmd

import (
	"fmt"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
)

var (
	ipRange string

	rootCmd = &cobra.Command{
		Use:   "nmapprom",
		Short: "Hosts connected to network",
		Long:  `Prometheus exporter for hosts connected to the network using Nmap.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := runServer(); err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	// Setup the root command with ip-range flag.
	rootCmd.PersistentFlags().StringVarP(&ipRange, "ip-range", "i", "", "IP range to scan")
	err := rootCmd.MarkPersistentFlagRequired("ip-range")
	if err != nil {
		panic(err)
	}
}

func Execute() error {
	return rootCmd.Execute()
}

func runServer() error {
	// Setup the chi http router.
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Initialise the 'hosts' variable to -1 - this represents the connected hosts to the network.
	hosts := -1
	// nmapChan is a channel that will be used to pass the number of hosts up to the chi http server.
	nmapChan := make(chan int)

	// Run the nmap command in a goroutine.
	go func() {
		for {
			runNmapChan(&nmapChan)
			time.Sleep(time.Minute)
		}
	}()

	// Fetch the number of hosts through the channel.
	go func() {
		for {
			hosts = <-nmapChan
		}
	}()

	// Set up the metrics endpoint to return the number of hosts.
	r.Get("/metrics", func(w http.ResponseWriter, r *http.Request) {
		// Return a 500 error if the number of hosts is -1 - indicating the nmap command hasn't run for the first time yet.
		if hosts == -1 {
			w.WriteHeader(500)
			return
		}
		fmt.Fprintf(w, "hosts_connected_network %d", hosts-1)
	})

	// Serve on port 3800
	err := http.ListenAndServe(":3800", r)
	if err != nil {
		return err
	}

	return nil
}

// Run nmap and pass the output into a channel.
func runNmapChan(channel *chan int) {
	// Run the nmap command.
	out, err := exec.Command("nmap", "-sn", "192.168.0.0/24").Output()
	if err != nil {
		panic(err)
	}

	// Parse the output to get the number of hosts.
	m := regexp.MustCompile(`\d+ hosts up`)
	res := strings.Split(m.FindStringSubmatch(string(out))[0], " ")
	in, err := strconv.Atoi(res[0])
	if err != nil {
		panic(err)
	}

	// Pass the number of hosts up to the chi http server through the channel.
	*channel <- in
}
