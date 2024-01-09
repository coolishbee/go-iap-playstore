# go-iap-playstore

![](https://img.shields.io/badge/golang-1.19-blue.svg?style=flat)

go-iap-playstore searches the refund list via Playstore.
This repository is inspired by [go-iap](https://github.com/awa/go-iap)

# Installation
```
go get github.com/coolishbee/go-iap-playstore
```


# Usage

## Search Refund List(Default)

```go
import(
    "github.com/coolishbee/go-iap-playstore"
)

func main() {
    ctx := context.Background()

    jsonFile, err := os.ReadFile("service-account.json")
	if err != nil {
		fmt.Println(err)
	}

	client, err := playstore.New(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	result, err := client.VoidedList(ctx, "packageName")
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
```

## Refund List search by time range

```go
import(
    "github.com/coolishbee/go-iap-playstore"
)

func main() {
	ctx := context.Background()
	
	jsonFile, err := os.ReadFile("service-account.json")
	if err != nil {
		fmt.Println(err)
	}

	client, err := playstore.New(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	endTime := time.Now().UTC()
	startTime := endTime.Add(time.Hour * -4)	
	result, err := client.VoidedListTimeRange(ctx, "packageName", startTime.UnixMilli(), endTime.UnixMilli())
	if err != nil {
		fmt.Println(err)
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
```