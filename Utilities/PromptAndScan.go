package Utilities

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptAndScan(prompt string, dest *string) error {
	_, err := fmt.Print(prompt)

	if err != nil {
		return err
	}
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	*dest = strings.TrimSpace(input)
	return nil
}
