package main

import (
	"flag"
	"fmt"

	"github.com/ItsArul/dockauto/cmd"
)

var (
	cli                                                           = cmd.GetConnect()
	id, port, newName, name, password, dockerfile, user, img, tag string
	help                                                          bool
)

func main() {

	id = *flag.String("id", "", "digunakan untuk inputan id dari id container ataupun image")
	port = *flag.String("port", "", "digunakan untuk binding port yang ingin dipublish")
	name = *flag.String("name", "", "digunakan untuk penamaan image ataupun container")
	password = *flag.String("password", "", "digunakan untuk auth ke dockerhub")
	user = *flag.String("user", "", "digunakan untuk user auth ke dockerhub")
	dockerfile = *flag.String("dockerfile", "", "digunakan untuk dockerfile")
	img = *flag.String("img", "", "digunakan untuk melakukan create container")
	newName = *flag.String("newname", "", "digunakan untuk rename container")
	tag = *flag.String("tag", "", "digunakan untuk tag dari image")
	help = *flag.Bool("help", false, "to know about infromation")

	flag.Parse()

	if help {
		flag.PrintDefaults()
	} else {
		Run()
	}
}

func Run() {
	values := flag.Args()
	if len(values) == 0 {
		fmt.Println("use --help to know about information")
	}

	for _, operation := range values {

		switch operation {
		case "container create":
			cmd.ContainerCreate(cli, img, port, name)
		case "container prune":
			cmd.ContainerPrune(cli)
		case "container logs":
			cmd.ContainerLogs(cli, id)
		case "container stop":
			cmd.ContainerStop(cli, id)
		case "container rename":
			cmd.ContaineRename(cli, id, newName)
		case "container list":
			cmd.ListContainer(cli)
		case "image build":
			cmd.ImageBuild(cli, dockerfile)
		case "image push":
			cmd.ImagePush(cli, tag, user, password)
		case "image list":
			cmd.ImageList(cli)
		case "image prune":
			cmd.ImagePrune(cli)
		case "image pull":
			cmd.ImagePull(cli, name)
		default:
			flag.PrintDefaults()
		}
	}
}
