package jobcan

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func GetStatus(email, password string) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var status string
	err := chromedp.Run(ctx,
		login(email, password),
		chromedp.Text("#working_status", &status, chromedp.NodeVisible, chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}

	return status, nil
}
