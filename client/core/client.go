package core

import (
	"fmt"
	"net"
)

type Client struct {
	Network string
	Address string
	Port    string
	Conn    net.Conn
}

func NewClient(network, address, port string) *Client {
	conn, err := net.Dial(network, address+":"+port)
	if err != nil {
		fmt.Printf("Create new client failed, err: %s\n", err.Error())
		return nil
	}
	return &Client{
		Network: network,
		Address: address,
		Port:    port,
		Conn:    conn,
	}

}

func (c *Client) ReadMessage() {
	for {

	}
}

func (c *Client) SendMessage() {

}
