package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type images struct {
	Subjects []imageLevel
}

type imageLevel struct {
	path string `json:"Id"`
}

func jsonParse(str []byte) ([]string ){
	// for test
//	str := []byte(`[
//    {
//        "Id": "sha256:16508e5c265dcb5c05017a2a8a8228ae12b7b56b2cda0197ed5411bda200a961",
//        "RepoTags": [
//            "ubuntu:latest"
//        ],
//        "RepoDigests": [
//            "ubuntu@sha256:ac533e4ead4110211a4d67cbf44ed8b7d1aca2b8e6f15d1e8768eadaf433dd31"
//        ],
//        "Parent": "",
//        "Comment": "",
//        "Created": "2018-08-22T17:28:57.797346492Z",
//        "Container": "9c1141adf35e335d53c7c5c051dfd56315201ec99f79310ff9099d19ee2eac15",
//        "ContainerConfig": {
//            "Hostname": "9c1141adf35e",
//            "Domainname": "",
//            "User": "",
//            "AttachStdin": false,
//            "AttachStdout": false,
//            "AttachStderr": false,
//            "Tty": false,
//            "OpenStdin": false,
//            "StdinOnce": false,
//            "Env": [
//                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
//            ],
//            "Cmd": [
//                "/bin/sh",
//                "-c",
//                "#(nop) ",
//                "CMD [\"/bin/bash\"]"
//            ],
//            "ArgsEscaped": true,
//            "Image": "sha256:4595a78b95f7994dfc56d6261c65ea7c5f2e69136484e5db30db34542c9ebc2d",
//            "Volumes": null,
//            "WorkingDir": "",
//            "Entrypoint": null,
//            "OnBuild": null,
//            "Labels": {}
//        },
//        "DockerVersion": "17.06.2-ce",
//        "Author": "",
//        "Config": {
//            "Hostname": "",
//            "Domainname": "",
//            "User": "",
//            "AttachStdin": false,
//            "AttachStdout": false,
//            "AttachStderr": false,
//            "Tty": false,
//            "OpenStdin": false,
//            "StdinOnce": false,
//            "Env": [
//                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
//            ],
//            "Cmd": [
//                "/bin/bash"
//            ],
//            "ArgsEscaped": true,
//            "Image": "sha256:4595a78b95f7994dfc56d6261c65ea7c5f2e69136484e5db30db34542c9ebc2d",
//            "Volumes": null,
//            "WorkingDir": "",
//            "Entrypoint": null,
//            "OnBuild": null,
//            "Labels": null
//        },
//        "Architecture": "amd64",
//        "Os": "linux",
//        "Size": 84117621,
//        "VirtualSize": 84117621,
//        "GraphDriver": {
//            "Data": {
//                "LowerDir": "/var/lib/docker/overlay2/bde2e51eb6a6ff63e1ccf1f37c7a30a0e4adc00c7fe289ef0fb060ef14725dd8/diff:/var/lib/docker/overlay2/240f9bc0ba505df075106d975e1a8f1283ac263ad9652d5c0f19d4abf07f8d65/diff:/var/lib/docker/overlay2/816f089cde039a608af06b57246c5a6e58ad6285a70da41522800816ef25e062/diff:/var/lib/docker/overlay2/211ba13f5fad41d407d1db31a0e2f07d4d04c3caaa3384b65d6eacce95d43109/diff",
//                "MergedDir": "/var/lib/docker/overlay2/3a049aa71370b564238744d50ff180761388592bd49a3088c826af3e1b96d4ad/merged",
//                "UpperDir": "/var/lib/docker/overlay2/3a049aa71370b564238744d50ff180761388592bd49a3088c826af3e1b96d4ad/diff",
//                "WorkDir": "/var/lib/docker/overlay2/3a049aa71370b564238744d50ff180761388592bd49a3088c826af3e1b96d4ad/work"
//            },
//            "Name": "overlay2"
//        },
//        "RootFS": {
//            "Type": "layers",
//            "Layers": [
//                "sha256:a30b835850bfd4c7e9495edf7085cedfad918219227c7157ff71e8afe2661f63",
//                "sha256:a2672464542154dc3fa4b31689055fe57dead430380899393285e4a190d042fe",
//                "sha256:b6a02001ba3318a8475fb7f54d7d870e4d7d71562b890b18d3cb58fe95bc6085",
//                "sha256:7422efa72a14c6f15af199905829c3d3fc3f7cadd7046e0ff1d2349ca6dfb159",
//                "sha256:ec8257ff6a7a399ac79e34864ce4e99b378c293f6e2865ea940f25df7982f34b"
//            ]
//        },
//        "Metadata": {
//            "LastTagTime": "0001-01-01T00:00:00Z"
//        }
//    }]
//`)
	//var out imageLevel
	//err := json.Unmarshal(str, &out)
	//if err != nil {
	//	fmt.Println("json err:", err)
	//}
	//fmt.Println(out)
	var t interface{}
	err := json.Unmarshal(str, &t)
	if err != nil {
		fmt.Println("json err:", err)
	}
	//fmt.Println(t)
	m := t.(map[string]interface{})
	//fmt.Println(m)
	g := m["GraphDriver"]
	//fmt.Println(g)
	h := g.(map[string]interface{})
	l := (h["Data"]).(map[string]interface{})
	//fmt.Println(l["LowerDir"])
	upperDir := l["UpperDir"]

	//fmt.Println(upperDir)
	//fmt.Println(reflect.TypeOf(upperDir))
	sss := "/var/lib/docker/overlay2/bde2e51eb6a6ff63e1ccf1f37c7a30a0e4adc00c7fe289ef0fb060ef14725dd8/diff:/var/lib/docker/overlay2/240f9bc0ba505df075106d975e1a8f1283ac263ad9652d5c0f19d4abf07f8d65/diff:/var/lib/docker/overlay2/816f089cde039a608af06b57246c5a6e58ad6285a70da41522800816ef25e062/diff:/var/lib/docker/overlay2/211ba13f5fad41d407d1db31a0e2f07d4d04c3caaa3384b65d6eacce95d43109/diff"
	lowerDirs := strings.Split(sss, ":")
	dirs := append(lowerDirs, upperDir.(string))

	//fmt.Println(dirs)
	//m = m["Metadata"].(map[string]interface{})
	//m = m["GraphDriver"].(map[string]interface{})
	//m = m["MergedDir"].(map[string]interface{})
	//fmt.Println(m)
	return dirs
}