package jobcan

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func LeaveWork(email, password string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err := chromedp.Run(ctx,
		login(email, password),
		chromedp.Click(`//*[@id="adit-button-work-end"]`, chromedp.NodeVisible),
		chromedp.Sleep(3*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
