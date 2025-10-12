package jsonkitty

import (
	"bytes"
	"encoding/json"
)

// Pretty преобразует строку JSON в отформатированный вид с отступами (4 пробела).
// Возвращает:
//   - []byte: Отформатированный JSON.
//   - error: Ошибка, если входные данные не являются корректным JSON.
func Pretty(jsonContent string) ([]byte, error) {
	var prettyJSON bytes.Buffer

	if err := json.Indent(&prettyJSON, []byte(jsonContent), "", "    "); err != nil {
		return []byte(""), err
	}

	return prettyJSON.Bytes(), nil
}
