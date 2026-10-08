// 17. Напиши функцию getWithRetry(url string): она
// отправляет GET-запрос клиентом с таймаутом 3 секунды,
// делает максимум 3 попытки с паузой 1 секунда между
// ними. Повторяет попытку при ошибке запроса или статусе
// 5xx; при статусе 4xx и 2xx сразу завершается без
// повторов. В main вызови функцию для
// https://httpbin.org/status/200,
// https://httpbin.org/status/404 и
// https://httpbin.org/status/503 и для каждого выведи
// итоговый статус-код и количество сделанных попыток.

package main

import (
	"fmt"
	"net/http"
	"time"
)

const maxAttempts = 3

// getWithRetry returns the final status code and the
// number of attempts. Status is 0 if no response came.
func getWithRetry(url string) (int, int) {
	client := &http.Client{Timeout: 3 * time.Second}

	status := 0
	attempts := 0

	for attempts < maxAttempts {
		attempts++

		resp, err := client.Get(url)
		if err != nil {
			status = 0
			fmt.Println("  request error:", err)
		} else {
			status = resp.StatusCode
			resp.Body.Close()

			// 2xx and 4xx: no retry
			if status < 500 {
				return status, attempts
			}
		}

		// pause only if another attempt follows
		if attempts < maxAttempts {
			time.Sleep(1 * time.Second)
		}
	}

	return status, attempts
}

func main() {
	urls := []string{
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/404",
		"https://httpbin.org/status/503",
	}

	for _, url := range urls {
		fmt.Println(url)
		status, attempts := getWithRetry(url)
		fmt.Printf("  status: %d, attempts: %d\n\n",
			status, attempts)
	}
}
