# SpotWorkload
Spot-Workload Representation and Management Component of the Technion's SpotOS Project

This repository contains the API for SpotWorkload CRD and the set of its managing controllers. 

## Getting Started

### SpotWorkload CR
A SpotWorkload CR defines a spot instance configuration as outputted from the [CloudCostOptimizer](https://github.com/AdiY10/CloudCostOptimizer), and reflects its utilization in status when the spot-instance begins serving as a cluster node.

### SpotIstance Controllers
The controllers are responsible for:
- Ordering the specified spot instance
- Installing the acquired spot instance as a cluster node
    - Including tainting and labeling properly
- Updating CR status with assigned node name

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

1.  Run the following command to deploy the `spot-instance-controller` controller to your cluster:
    ```
    envsubst < deploy/spot-instance-controller.yaml.template | kubectl apply -f -
    ```

## Cleanup from a cluster

1.  Run the following command to clean `spot-instance-controller` from your leaf hub cluster:
    ```
    envsubst < deploy/spot-instance-controller.yaml.template | kubectl delete -f -
    ```
