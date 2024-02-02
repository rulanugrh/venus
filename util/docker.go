package util

import docker "github.com/fsouza/go-dockerclient"

func GetClient() *docker.Client {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		return nil
	}

	return client
}