package price

import (
	"fmt"
	"github.com/evleria-trading/position-client/internal/scope"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

func NewListCmd(s *scope.Scope) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(s)
		},
	}

	return cmd
}

func runList(s *scope.Scope) error {
	prices := s.PricesCache.GetPrices()
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintf(w, "Symbol\tAsk\tBid\t\n")
	for _, pr := range prices {
		fmt.Fprintf(w, "%s\t%.2f\t%.2f\t\n", pr.Symbol, pr.Ask, pr.Bid)
	}
	w.Flush()

	return nil
}
