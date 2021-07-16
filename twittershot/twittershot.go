package twittershot

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

func Screenshot(path, sel string) ([]byte, error) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx, eleScreenshot(path, sel, &buf)); err != nil {
		return buf, err
	}
	return buf, nil
}

func eleScreenshot(path, sel string, buf *[]byte) chromedp.Tasks {
	width, height := 1920, 1080
	var res string
	return chromedp.Tasks{
		emulation.SetDeviceMetricsOverride(int64(width), int64(height), 1.0, false),
		chromedp.Navigate(path),
		chromedp.Sleep(time.Second * 2),
		chromedp.EvaluateAsDevTools("window.scroll(0,0); '';", &res),
		chromedp.Sleep(time.Second * 2),
		chromedp.Screenshot(sel, buf, chromedp.NodeVisible),
	}
}
