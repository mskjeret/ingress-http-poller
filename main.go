package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	var kubeconfig *string

	home := homeDir()
	folder := filepath.Join(home, ".kube")
	path := filepath.Join(folder, "config")
	writeKubeConfigToDiskFromEnvironment(folder, path)
	kubeconfig = flag.String("kubeconfig", path, "(optional) absolute path to the kubeconfig file")

	flag.Parse()

	urls := ResolveIngress(kubeconfig)
	fmt.Printf("Found %d hosts to check:\n\n", len(urls))

	var working strings.Builder
	var notWorking strings.Builder

	for _, url := range urls {
		var responseCode = ExecuteURL(url)
		if responseCode > 199 && responseCode < 400 {
			working.WriteString(fmt.Sprintf("Executing %q and got response code %d\n", url, responseCode))
		} else {
			notWorking.WriteString(fmt.Sprintf("Executing %q and got response code %d\n", url, responseCode))
		}
	}

	var resultOutput strings.Builder
	if notWorking.Len() > 0 {

		resultOutput.WriteString(":dizzy_face: Following urls do not respond correctly:\n\n")
		resultOutput.WriteString(notWorking.String())
		resultOutput.WriteString("\n\nHowever, following urls do seem to work:\n\n")
		resultOutput.WriteString(working.String())
		NotifySlack(getSlackChannel(), getSlackAPIKey(), resultOutput.String())
	} else {
		resultOutput.WriteString("All hosts seems to be working")
	}

	fmt.Print(resultOutput.String())
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func writeKubeConfigToDiskFromEnvironment(folder string, path string) {
	config := os.Getenv("KUBECONFIG")
	if config != "" {
		os.MkdirAll(folder, os.ModePerm)
		message := []byte(config)
		err := ioutil.WriteFile(path, message, 0644)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		panic("KUBECONFIG key not set")
	}
}

func getSlackAPIKey() string {
	key := os.Getenv("SLACK_API_KEY")
	if key == "" {
		panic("SLACK_API_KEY not set")
	}
	return key
}

func getSlackChannel() string {
	key := os.Getenv("SLACK_CHANNEL")
	if key == "" {
		panic("SLACK_CHANNEL not set")
	}
	return key
}
