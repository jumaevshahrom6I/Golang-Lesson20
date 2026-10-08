// 9. Для кодов 200, 201, 404, 500, 503 запроси
// https://httpbin.org/status/<код> клиентом с таймаутом
// 5 секунд. Для каждого выведи "<код>: успех",
// "<код>: ошибка клиента" или "<код>: ошибка сервера".

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{Timeout: 5 * time.Second}
	codes := []int{200, 201, 404, 500, 503}

	for _, code := range codes {
		url := fmt.Sprintf("https://httpbin.org/status/%d", code)

		resp, err := client.Get(url)
		if err != nil {
			fmt.Printf("%d: request error: %v\n", code, err)
			continue
		}
		status := resp.StatusCode
		resp.Body.Close()

		switch {
		case status >= 200 && status < 300:
			fmt.Printf("%d: success\n", status)
		case status >= 400 && status < 500:
			fmt.Printf("%d: client error\n", status)
		case status >= 500:
			fmt.Printf("%d: server error\n", status)
		default:
			fmt.Printf("%d: other status\n", status)
		}
	}
}
