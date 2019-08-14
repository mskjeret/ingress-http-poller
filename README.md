# Simple uptime monitor for you k8s ingress

This project will try to keep an eye on you ingresses in the given kubernetes cluster and check that they are indeed responding on the internet.
It can be run standalone or as a docker container. 

Currently it requires a Slack integration and hence need an API key and a channel to send to.

## The kube config

The kube config must be set as the environment variable KUBECONFIG

## Slack integration

Set the API key in the environment variable SLACK_API_KEY and then set the channel to send to in the environment variable SLACK_CHANNEL

## Docker command

docker run --rm -e KUBECONFIG=the_config -e SLACK_API_KEY=the_key -e SLACK_CHANNEL=channel mskjeret/ingress-http-poller:latest



