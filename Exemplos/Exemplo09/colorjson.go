/* https://github.com/TylerBrock/colorjson */
package main

import (
	"encoding/json"
	"fmt"

	"github.com/TylerBrock/colorjson"
)

func main() {
	str := `{
      "cpf": "123.123.123-12",
      "num": 100,
      "bool": false,
      "null": null,
      "array": ["foo", "bar", "baz"],
      "obj": { "a": 1, "b": 2 }
    }`

	var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)

	// Make a custom formatter with indent set
	f := colorjson.NewFormatter()
	f.Indent = 4

	// Marshall the Colorized JSON
	s, _ := f.Marshal(obj)
	fmt.Println(string(s))
}
