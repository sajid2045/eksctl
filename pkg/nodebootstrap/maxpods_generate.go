//+build ignore

package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"net/http"

	. "github.com/dave/jennifer/jen"
)

const maxPodsPerNodeTypeSourceText = "https://raw.github.com/awslabs/amazon-eks-ami/master/files/eni-max-pods.txt"

func main() {
	f := NewFile("nodebootstrap")

	f.Comment("Generated from " + maxPodsPerNodeTypeSourceText)

	d := Dict{}

	resp, err := http.Get(maxPodsPerNodeTypeSourceText)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, line := range strings.Split(string(body), "\n") {
		if !strings.HasPrefix(line, "#") {
			parts := strings.Split(line, " ")
			if len(parts) == 2 {
				k := Lit(parts[0])
				v, err := strconv.Atoi(parts[1])
				if err != nil {
					log.Fatal(err.Error())
				}
				d[k] = Lit(v)
			}
		}
	}

	f.Var().Id("maxPodsPerNodeType").Op("=").
		Map(String()).Int().Values(d)

	if err := f.Save("maxpods.go"); err != nil {
		log.Fatal(err.Error())
	}
}
