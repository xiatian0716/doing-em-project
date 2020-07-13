package utils

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"go-crawler-test/config"
	"log"
	"time"
)

var ctxRateLimiter = time.Tick(config.RateLimiter)

func OptionsUtil() []chromedp.ExecAllocatorOption {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	return options
}

// setcookies returns a task to navigate to a host with the passed cookies set
// on the network request.
func HeadersUtil(url string, debug bool, domain string, headers map[string]interface{}, cookies ...string) chromedp.Tasks {
	<-ctxRateLimiter
	log.Printf("Fetching url:%s", url)
	if len(cookies)%2 != 0 {
		panic("length of cookies must be divisible by 2")
	}
	return chromedp.Tasks{
		// add headers to chrome
		network.Enable(),
		network.SetExtraHTTPHeaders(network.Headers(headers)),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// create cookie expiration
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			// add cookies to chrome
			for i := 0; i < len(cookies); i += 2 {
				success, err := network.SetCookie(cookies[i], cookies[i+1]).
					WithExpires(&expr).
					WithDomain(domain).
					WithHTTPOnly(true).
					Do(ctx)
				if err != nil {
					return err
				}
				if !success {
					return fmt.Errorf("could not set cookie %q to %q", cookies[i], cookies[i+1])
				}
			}
			return nil
		}),
		// navigate to site
		chromedp.Navigate(url),
		// read network values
		chromedp.ActionFunc(func(ctx context.Context) error {
			if debug {
				cookies, err := network.GetAllCookies().Do(ctx)
				if err != nil {
					return err
				}
				for i, cookie := range cookies {
					log.Printf("chrome cookie %d: %+v", i, cookie)
				}
			}
			return nil
		}),
	}
}
