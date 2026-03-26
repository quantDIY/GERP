package cli

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/spanner"
	"github.com/spf13/cobra"
	"google.golang.org/api/iterator"
)

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Deep Spanner inspection for Legal/Compliance bounds",
}

var viewCmd = &cobra.Command{
	Use:   "view [target_record_id]",
	Short: "Bypasses the BFF to directly query Spanner ComplianceAudits locally",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		recordID := args[0]
		fmt.Printf("🔍 Establishing secure isolation boundaries to Spanner: %s...\n", recordID)

		ctx := context.Background()
		client, err := spanner.NewClient(ctx, ActiveConfig.SpannerDB)
		if err != nil {
			fmt.Println("🚨 CRITICAL: Sovereign Spanner binding failed:", err)
			os.Exit(1)
		}
		defer client.Close()

		stmt := spanner.Statement{
			SQL: `SELECT ID, ActorID, Action, AuditTimestamp 
			      FROM ComplianceAudits 
			      WHERE TargetRecordID = @record_id`,
			Params: map[string]interface{}{
				"record_id": recordID,
			},
		}

		iter := client.Single().Query(ctx, stmt)
		defer iter.Stop()

		fmt.Println("\nImmutable Ledger Responses:")
		fmt.Println("==================================================")
		found := false

		// Bypassing typed structs strictly for administrative debug terminal printing
		for {
			row, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				fmt.Println("🚨 CRITICAL: Matrix iteration failure:", err)
				os.Exit(1)
			}
			found = true
			
			var id, actorID, action string
			var auditTimestamp spanner.NullTime

			if err := row.Columns(&id, &actorID, &action, &auditTimestamp); err != nil {
				fmt.Println("🚨 Error mapping compliance row:", err)
				continue
			}

			fmt.Printf("Audit ID:   %s\n", id)
			fmt.Printf("Action:     %s\n", action)
			fmt.Printf("Actor:      %s\n", actorID)
			fmt.Printf("Timestamp:  %v\n", auditTimestamp.Time)
			fmt.Println("--------------------------------------------------")
		}

		if !found {
			fmt.Println("Warning: Target record un-auditable natively. Zero matrix matches bound.")
		}
	},
}

func init() {
	auditCmd.AddCommand(viewCmd)
	rootCmd.AddCommand(auditCmd)
}
