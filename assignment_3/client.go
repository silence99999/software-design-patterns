package assignment_3

import "fmt"

type Client struct{}

func (c *Client) PrintText(file File) {
	fmt.Println("Client sends the file")
	file.PrintInJSON()
}
