package voided_purchases

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestVoidedList(t *testing.T) {
	ctx := context.Background()

	jsonFile, err := os.ReadFile("service-account.json")
	if err != nil {
		t.Errorf("%s", err)
	}

	client, err := New(jsonFile)
	if err != nil {
		t.Errorf("got %#v", err)
	}
	result, err := client.VoidedList(ctx, "com.gamepub.hab.g")
	if err != nil {
		t.Errorf("got %v", err)
	}
	if result != nil {
		if result.TokenPagination != nil {
			fmt.Println(result.TokenPagination.NextPageToken)
		} else {
			fmt.Println("result.TokenPagination nil")
		}
		if result.VoidedPurchases != nil {
			for _, item := range result.VoidedPurchases {
				fmt.Println(item.PurchaseToken)
				fmt.Println(item.PurchaseTimeMillis)
				fmt.Println(item.VoidedTimeMillis)
				fmt.Println(item.OrderId)
				fmt.Println(item.VoidedSource)
				fmt.Println(item.VoidedReason)
				fmt.Println(item.Kind)
			}
		} else {
			fmt.Println("result.VoidedPurchases nil")
		}
	} else {
		fmt.Println("result nil")
	}
}
