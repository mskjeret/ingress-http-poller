package main

import (
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//ResolveIngress Resolve all the ingress and paths to test
func ResolveIngress(kubeconfig *string) []string {
	var urls []string

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	ingressList, err := clientset.ExtensionsV1beta1().Ingresses("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for i := 0; i < len(ingressList.Items); i++ {
		ingress := ingressList.Items[i]
		rules := ingress.Spec.Rules
		tlsArray := ingress.Spec.TLS

		var secureHosts []string
		for t := 0; t < len(tlsArray); t++ {
			tls := tlsArray[t]
			hosts := tls.Hosts
			for _, s := range hosts {
				secureHosts = append(secureHosts, s)
			}
		}

		for j := 0; j < len(rules); j++ {
			rule := rules[j]
			paths := rule.HTTP.Paths
			isSSL := contains(secureHosts, rule.Host)

			var buffer strings.Builder
			if isSSL {
				buffer.WriteString("https://")
			} else {
				buffer.WriteString("http://")
			}
			buffer.WriteString(rule.Host)
			for _, s := range paths {
				urls = append(urls, buffer.String()+s.Path)
			}
		}
	}

	return urls
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
