package cmd

import "github.com/docker/docker/client"

func GetConnect() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	return cli

}
