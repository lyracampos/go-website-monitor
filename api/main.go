package main

import (
	"errors"
	"fmt"
)

func main() {
	metadata := getMetadado()

	if item, exists := metadata["error"]; exists {
		xtype := fmt.Sprintf("%T", item)
		fmt.Println(xtype)
	}

	if item, exists := metadata["int"]; exists {
		xtype := fmt.Sprintf("%T", item)
		fmt.Println(xtype)
	}

	if item, exists := metadata["bool"]; exists {
		xtype := fmt.Sprintf("%T", item)
		fmt.Println(xtype)
	}
}

func getMetadado() map[string]interface{} {
	metadata := map[string]interface{}{
		"int":   1,
		"error": errors.New("error tal"),
		"bool":  false,
	}
	return metadata
}
