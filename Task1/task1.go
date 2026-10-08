// 1. Отправь GET-запрос на
// https://jsonplaceholder.typicode.com/posts/1. Выведи
// статус-код, значение заголовка ответа Content-Type и
// тело ответа как строку.

package main

import (
	"fmt"
	"io"
	"net/http"
)

const base = "https://jsonplaceholder.typicode.com"

func main() {
	resp, err := http.Get(base + "/posts/1")
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

	fmt.Println("Status code:", resp.StatusCode)
	fmt.Println("Content-Type:", resp.Header.Get("Content-Type"))
	fmt.Println("Body:", string(body))
}
