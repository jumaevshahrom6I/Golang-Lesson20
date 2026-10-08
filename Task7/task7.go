// 7. Отправь GET-запрос на https://httpbin.org/headers,
// добавив в запрос заголовок X-Student со своим именем.
// Выведи тело ответа и убедись, что твой заголовок в
// нём есть.

package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	req, err := http.NewRequest(
		http.MethodGet, "https://httpbin.org/headers", nil)
	if err != nil {
		fmt.Println("Request creation error:", err)
		return
	}
	req.Header.Set("X-Student", "Shahrom")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Println(string(body))

	if strings.Contains(string(body), "X-Student") {
		fmt.Println("X-Student header found in response")
	} else {
		fmt.Println("X-Student header not found")
	}
}
