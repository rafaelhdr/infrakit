package loadbalancer

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/docker/infrakit/pkg/cli"
	"github.com/docker/infrakit/pkg/spi/loadbalancer"
	"github.com/spf13/cobra"
)

// Routes returns the describe command
func Routes(name string, services *cli.Services) *cobra.Command {
	routes := &cobra.Command{
		Use:   "routes",
		Short: "Loadbalancer routes",
	}

	ls := &cobra.Command{
		Use:   "ls",
		Short: "List loadbalancer routes",
	}

	publish := &cobra.Command{
		Use:   "add",
		Short: "Publish a loadbalancer route",
	}
	unpublish := &cobra.Command{
		Use:   "rm <frontend_port>, ....",
		Short: "Unpublish a loadbalancer route",
	}

	routes.AddCommand(ls, publish, unpublish)

	port := publish.Flags().Int("port", 80, "Backend listening port")
	protocol := publish.Flags().String("protocol", "http", "Protocol: http|https|tcp|udp|ssl")
	frontendPort := publish.Flags().Int("frontend-port", 80, "Frontend loadbalancer port")
	cert := publish.Flags().String("cert", "", "Certificate")

	publish.RunE = func(cmd *cobra.Command, args []string) error {
		l4, err := Load(services.Plugins(), name)
		if err != nil {
			return nil
		}
		cli.MustNotNil(l4, "L4 not found", "name", name)

		route := loadbalancer.Route{
			Port:             *port,
			LoadBalancerPort: *frontendPort,
			Protocol:         loadbalancer.Protocol(*protocol),
		}
		if *cert != "" {
			copy := *cert
			route.Certificate = &copy
		}
		res, err := l4.Publish(route)
		fmt.Println(res)
		return err
	}

	unpublish.RunE = func(cmd *cobra.Command, args []string) error {
		l4, err := Load(services.Plugins(), name)
		if err != nil {
			return nil
		}
		cli.MustNotNil(l4, "L4 not found", "name", name)

		targets := []int{}
		for _, a := range args {
			v, err := strconv.Atoi(a)
			if err != nil {
				return err
			}
			targets = append(targets, v)
		}

		for _, t := range targets {
			res, err := l4.Unpublish(t)
			if err != nil {
				return err
			}
			fmt.Println(res)
		}
		return err
	}

	ls.Flags().AddFlagSet(services.OutputFlags)
	ls.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			cmd.Usage()
			os.Exit(1)
		}

		l4, err := Load(services.Plugins(), name)
		if err != nil {
			return nil
		}
		cli.MustNotNil(l4, "L4 not found", "name", name)

		list, err := l4.Routes()
		if err != nil {
			return err
		}

		return services.Output(os.Stdout, list,
			func(w io.Writer, v interface{}) error {

				fmt.Printf("%-15v  %-10v %-10v  %-20v\n", "FRONTEND PORT", "PROTOCOL", "BACKEND PORT", "CERT")
				for _, r := range list {
					cert := ""
					if r.Certificate != nil {
						cert = *r.Certificate
					}
					fmt.Printf("%-15v  %-10v %-10v %-20v\n", r.LoadBalancerPort, r.Protocol, r.Port, cert)
				}
				return nil
			})
	}
	return routes
}
