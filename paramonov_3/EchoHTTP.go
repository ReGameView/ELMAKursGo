package paramonov_3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PrettyJson struct {
}

// startEchoHandler -- Создание эхо сервера
func EchoHandler(port string) {

	http.HandleFunc("/", Handler)

	err := http.ListenAndServe(":"+port, nil)

	fmt.Println("Ошибка при создание сервера: ", err)
}

// EchoHandler Эхо возращает POST ответ
func Handler(writer http.ResponseWriter, request *http.Request) {
	var header string = request.Header.Get("Content-Type")

	body, _ := ioutil.ReadAll(request.Body)

	_, _ = writer.Write([]byte("Header: " + header + "\n"))
	_, _ = writer.Write([]byte("Body\n" + string(body) + "\n"))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

}
