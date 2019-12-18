package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type Container struct {
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	check(err)

	// https://kuroeveryday.blogspot.com/2017/09/golang-build-image-with-dockerfile.html
	// https://qiita.com/nkz0914ssk/items/5429e75e5c0711add93a
	//https://github.com/dbhagen/sidefx-hkey-docker
	//https://github.com/adamrehn/ue4-cloud-rendering-demo

	// RUN image
	//reader, err := cli.ImagePull(ctx, "docker.io/library/centos", types.ImagePullOptions{})
	//if err != nil {
	//	panic(err)
	//}
	//io.Copy(os.Stdout, reader)

	// Dockerfile.tar.gzを読み込む
	cwd, _ := os.Getwd()
	file, err := os.Open(cwd + "/Dockerfile.tar.gz")
	defer file.Close()

	// イメージ名はUnix時間
	imageName := fmt.Sprintf("%v", time.Now().Unix())

	// Dockerfileからイメージを作成する
	res, err := cli.ImageBuild(ctx, file, types.ImageBuildOptions{
		Tags:        []string{imageName},
		ForceRemove: true,
	})
	check(err)
	defer res.Body.Close()

	fmt.Printf("OSType: %s\n", res.OSType)

	b, err := ioutil.ReadAll(res.Body)
	check(err)
	fmt.Println(string(b))

	fmt.Printf("Build Image. Image's name is %v\n", imageName)

	// イメージ名を指定しコンテナを作成する
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
	}, nil, nil, "")
	check(err)

	// コンテナの実行
	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	check(err)

	_, err = cli.ContainerWait(ctx, resp.ID)
	check(err)

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{
		ShowStdout: true,
	})
	check(err)

	// 実行結果の出力
	io.Copy(os.Stdout, out)

	// 起動したコンテナを削除する
	err = cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{})
	check(err)

	//resp, err := cli.ContainerCreate(ctx, &container.Config{
	//	Image: "centos",
	//	Cmd:   []string{"echo", "hello"},
	//}, nil, nil, "")
	//if err != nil {
	//	panic(err)
	//}
	//
	//if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
	//	panic(err)
	//}
	//
	//// EXEC command
	//econfig := types.ExecConfig{AttachStdout: true, Cmd: []string{"echo", "aaa"}}
	//id, err := cli.ContainerExecCreate(ctx, resp.ID, econfig)
	//if err != nil {
	//	panic(err)
	//}
	//
	//res, err := cli.ContainerExecAttach(ctx, id.ID, econfig)
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer res.Close()
	//
	//stdout := new(bytes.Buffer)
	//stderr := new(bytes.Buffer)
	//_, err = stdcopy.StdCopy(stdout, stderr, res.Reader)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(stdout.String())
	//
	//
	//if _, err := cli.ContainerWait(ctx, resp.ID); err != nil {
	//	panic(err)
	//}
	//
	//// LOGGING
	//out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	//if err != nil {
	//	panic(err)
	//}
	//stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}
