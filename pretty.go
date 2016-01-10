package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func PrettyJSON(s string) string {
	var pretty bytes.Buffer
	json.Indent(&pretty, []byte(s), "", "\t")

	return fmt.Sprintf("%s", string(pretty.Bytes()))
}
