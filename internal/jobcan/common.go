package jobcan

import "github.com/chromedp/chromedp"

func login(email, password string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate("https://id.jobcan.jp/users/sign_in?app_key=atd"),
		chromedp.SendKeys("#user_email", email, chromedp.NodeVisible, chromedp.ByID),
		chromedp.SendKeys("#user_password", password, chromedp.NodeVisible, chromedp.ByID),
		chromedp.Submit("#user_email"),
	}
}
