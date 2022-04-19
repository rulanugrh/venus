package cmd

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func ImagePull(cli *client.Client, name string) {
	image, err := cli.ImagePull(context.Background(), name, types.ImagePullOptions{})
	if err != nil {
		fmt.Println("cannot pull image")
		panic(err)
	}

	defer image.Close()
	io.Copy(os.Stdout, image)
}

func ImageList(cli *client.Client) {
	image, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		fmt.Println("cannot list image")
	}

	for _, img := range image {
		fmt.Printf("tags image %s\n", img.ID)
	}
}

func ImageBuild(cli *client.Client, dockerfile string) {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	dockefileReader, err := os.Open(dockerfile)
	if err != nil {
		fmt.Println("cannot load dockerfile")
		panic(err)
	}

	readdockerfile, errDocker := ioutil.ReadAll(dockefileReader)
	if err != nil {
		fmt.Println("cannot read dockerfile")
		panic(errDocker)
	}

	tarHeader := &tar.Header{
		Name: dockerfile,
		Size: int64(len(readdockerfile)),
	}

	errTar := tw.WriteHeader(tarHeader)
	if errTar != nil {
		panic(errTar)
	}

	_, errWrite := tw.Write(readdockerfile)
	if errWrite != nil {
		panic(errWrite)
	}

	dockerFileReaderTar := bytes.NewBuffer(buf.Bytes())

	buildOptions := types.ImageBuildOptions{
		Context:    dockerFileReaderTar,
		Dockerfile: dockerfile,
		Remove:     true,
	}

	imageBuild, errBuild := cli.ImageBuild(context.Background(), dockerFileReaderTar, buildOptions)
	if errBuild != nil {
		fmt.Println("cannot build dockerimage")
		panic(err)
	}
	defer imageBuild.Body.Close()

}

func ImagePrune(cli *client.Client) {
	imgprune, err := cli.ImagesPrune(context.Background(), filters.Args{})
	if err != nil {
		fmt.Println("cannot prune image")
		panic(err)
	}

	fmt.Printf("image prune %s\n", imgprune.ImagesDeleted)
}

func ImagePush(cli *client.Client, tag string, user string, password string) {

	authDocker := types.AuthConfig{
		Username: user,
		Password: password,
	}

	authJSON, _ := json.Marshal(authDocker)
	authEncoded := base64.URLEncoding.EncodeToString(authJSON)

	opts := types.ImagePushOptions{RegistryAuth: authEncoded}
	reader, err := cli.ImagePush(context.Background(), tag, opts)
	if err != nil {
		fmt.Println("cannot push dockerimage")
		panic(err)
	}

	defer reader.Close()
}
