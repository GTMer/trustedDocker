package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/cbergoon/merkletree"
	"log"
	"os"
	"os/exec"
)

func Cmd(cmd string, shell bool) []byte {
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			panic("some error found")
		}
		return out
	} else {
		out, err := exec.Command(cmd).Output()
		if err != nil {
			panic("some error found")
		}
		return out
	
	}
}

func main() {
	//str := "docker inspect"
	//arg := os.Args[1]
	//fmt.Println(reflect.TypeOf(str))
	//fmt.Println(reflect.TypeOf(arg))
	//fmt.Println(arg)
	//str = arg

	//out := Cmd("docker inspect" + arg, true)
	app := "docker"

	arg0 := "inspect"

	arg1 := os.Args[1]
        fmt.Println(arg1)
	cmd := exec.Command(app, arg0, arg1)
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}
	//fmt.Println("**")
	//fmt.Println(string(stdout[1:len(stdout) - 2]))
	//fmt.Println("**")
	//fmt.Println(reflect.TypeOf(stdout))
	dirs := jsonParse(stdout[1:len(stdout) - 2])

	fmt.Println(dirs)
	//fmt.Print(out)

	var hashs [] string
	for i := range dirs {
		temp := dirs[i] + "/*"
		fmt.Println(temp)
		//cmd := exec.Command("find", temp, "-type", "f", "-print0", "|" ,"sort" ,"-z", "|", "xargs", "-0" ,"sha3sum", "|", "sha3sum")
		//println(cmd.Args)
		//stdout, err := cmd.Output()
		//if err != nil {
		//	println(err.Error())
		//	return
		//}
		//hashs = append(hashs, string(stdout[:]))
		hashs = append(hashs, string(Cmd("find "+ temp + " -type f -print0 | sort -z | xargs -0 sha3sum | sha3sum", true))[:56])
		//fmt.Println()
	}

	fmt.Println(hashs)

	var list []merkletree.Content
	for i := range hashs {
		list = append(list, TestContent{x: hashs[i]})
	}

	//Create a new Merkle Tree from the list of Content
	t, err := merkletree.NewTree(list)
	if err != nil {
		log.Fatal(err)
	}

	//Get the Merkle Root of the tree
	mr := t.MerkleRoot()
	log.Println(mr)

	//Verify the entire tree (hashes for each node) is valid
	vt, err := t.VerifyTree()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Verify Tree: ", vt)

	//Verify a specific content in in the tree
	fmt.Println("***")
	remoteRoot := getRoot("http://localhost:8080")
	fmt.Println(remoteRoot)
	vc, err := t.VerifyContent(TestContent{x:remoteRoot})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(vc)
	if vc == true {
                Cmd("openstack appcontainer run --net network=b1c73230-95af-48f3-9891-7a4399bafd17 " + arg1, true)
		//Cmd("docker run -it " + arg1, true)
	}


}

//TestContent implements the Content interface provided by merkletree and represents the content stored in the tree.
type TestContent struct {
	x string
}

//CalculateHash hashes the values of a TestContent
func (t TestContent) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(t.x)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

//Equals tests for equality of two Contents
func (t TestContent) Equals(other merkletree.Content) (bool, error) {
	return t.x == other.(TestContent).x, nil
}
