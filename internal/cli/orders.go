package cli

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var ordersCmd = &cobra.Command{
	Use:   "orders",
	Short: "Manage revenue sales orders natively",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Triggers the Global Fulfillment Saga via GraphQL",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("⚡ Initiating Global Fulfillment Sequence via BFF...")

		// The raw GraphQL payload injected using the Genesis Mocks (Deterministic)
		payload := []byte(`{
			"query": "mutation ExecuteGlobalSale($input: CreateOrderInput!) { createSalesOrder(input: $input) { id status totalValue customer { legalName countryCode } } }",
			"variables": {
				"input": {
					"transactionId": "99999999-9999-9999-9999-999999999999",
					"customerId": "22222222-2222-2222-2222-222222222222",
					"accountDebitId": "77777777-7777-7777-7777-777777777777",
					"accountCreditId": "88888888-8888-8888-8888-888888888888",
					"totalAmountCents": 250000,
					"items": [
						{
							"lotId": "66666666-6666-6666-6666-666666666666",
							"quantity": 5
						}
					]
				}
			}
		}`)

		req, err := http.NewRequest("POST", ActiveConfig.GraphQLEndpoint, bytes.NewBuffer(payload))
		if err != nil {
			fmt.Println("🚨 CRITICAL: API formulation fault:", err)
			os.Exit(1)
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("🚨 CRITICAL: Matrix unreachable via BFF:", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("✅ Gateway Matrix Traversal Complete. BFF Payload:\n%s\n", string(body))
	},
}

func init() {
	ordersCmd.AddCommand(createCmd)
	rootCmd.AddCommand(ordersCmd)
}
