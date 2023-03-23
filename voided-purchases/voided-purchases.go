package voided_purchases

import (
	"context"
	"net/http"
	"time"

	"google.golang.org/api/option"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	androidpublisher "google.golang.org/api/androidpublisher/v3"
)

type Client struct {
	service *androidpublisher.Service
}

func New(jsonKey []byte) (*Client, error) {
	c := &http.Client{Timeout: 10 * time.Second}
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, c)

	conf, err := google.JWTConfigFromJSON(jsonKey, androidpublisher.AndroidpublisherScope)
	if err != nil {
		return nil, err
	}

	val := conf.Client(ctx).Transport.(*oauth2.Transport)
	_, err = val.Source.Token()
	if err != nil {
		return nil, err
	}

	service, err := androidpublisher.NewService(ctx, option.WithHTTPClient(conf.Client(ctx)))
	if err != nil {
		return nil, err
	}

	return &Client{service}, err
}

func (c *Client) VoidedList(
	ctx context.Context,
	packageName string,
) (*androidpublisher.VoidedPurchasesListResponse, error) {
	ps := androidpublisher.NewPurchasesVoidedpurchasesService(c.service)
	result, err := ps.List(packageName).Context(ctx).Do()

	return result, err
}
