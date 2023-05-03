# SpotWorkload
Spot-Workload Representation and Management Component of the Technion's SpotOS Project

This repository contains the API for SpotWorkload CRD and the set of its managing controllers. 

## Getting Started

### SpotWorkload CR
A SpotWorkload CR defines a workload configuration as expected to input in [CloudCostOptimizer](https://github.com/AdiY10/CloudCostOptimizer),
and reflects the workload's scheduling decisions and status as it progresses.

### SpotWorkload Controllers
The controllers are responsible for:
- Translating the workload's scheduling decisions into tolerations and node affinity rules

TODO: enhance docs, add examples

## Build and push the image to docker registry

1.  Set the `REGISTRY` environment variable to hold the name of your docker registry:
    ```
    $ export REGISTRY=...
    ```

1.  Set the `IMAGE_TAG` environment variable to hold the required version of the image.  
    default value is `latest`, so in that case no need to specify this variable:
    ```
    $ export IMAGE_TAG=latest
    ```

1.  Run make to build and push the image:
    ```
    $ make push-images
    ```

## Deploy on a cluster

1.  Set the `REGISTRY` environment variable to hold the name of your docker registry:
    ```
    $ export REGISTRY=...
    ```

1.  Set the `IMAGE` environment variable to hold the name of the image.

    ```
    $ export IMAGE=$REGISTRY/$(basename $(pwd)):latest
    ```

1.  Run the following command to deploy the `spot-workload-controller` controller to your cluster:
    ```
    envsubst < deploy/spot-workload-controller.yaml.template | kubectl apply -f -
    ```

## Cleanup from a cluster

1.  Run the following command to clean `spot-workload-controller` from your cluster:
    ```
    envsubst < deploy/spot-workload-controller.yaml.template | kubectl delete -f -
    ```
