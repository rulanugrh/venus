package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func ContainerCreate(cli *client.Client, image string, port string, name string) {
	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: port,
	}

	containerPort, errContainer := nat.NewPort("tcp", port)
	if errContainer != nil {
		fmt.Println("cant binding port")
		panic(errContainer)
	}

	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
	cont, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: image,
		},
		&container.HostConfig{
			PortBindings: portBinding,
		},
		nil, nil, name)

	if err != nil {
		fmt.Println("cant create container")
		panic(err)
	}

	cli.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})
	fmt.Printf("container %v is started ", cont.ID)
}

func ListContainer(cli *client.Client) {
	container, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	if len(container) > 0 {
		for _, containers := range container {
			fmt.Printf("contaainer id running %s", containers.ID)
		}
	} else {
		fmt.Println("no container running")
	}

}

func ContainerStop(cli *client.Client, containerId string) {
	err := cli.ContainerStop(context.Background(), containerId, nil)
	if err != nil {
		panic(err)
	}
}

func ContainerPrune(cli *client.Client) {
	container, err := cli.ContainersPrune(context.Background(), filters.Args{})
	if err != nil {
		fmt.Println("cannot prune container")
	}

	fmt.Printf("container prune %v\n", container.ContainersDeleted)
}

func ContaineRename(cli *client.Client, id string, newName string) {
	err := cli.ContainerRename(context.Background(), id, newName)
	if err != nil {
		panic(err)
	}
}

func ContainerLogs(cli *client.Client, id string) {
	cont, err := cli.ContainerLogs(context.Background(), id, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})

	if err != nil {
		fmt.Println("cannot find logs")
		panic(err)
	}

	bufferRead, err := io.ReadAll(cont)
	if err != nil {
		fmt.Println("cannot read logs")
		panic(err)
	}

	fmt.Println("container logs" + string(bufferRead))
}
