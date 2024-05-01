# Delve Debugger Demo with Kubernetes and Tilt

This repository demonstrates how to set up and use the Delve debugger inside a Docker container, orchestrated by Kubernetes, with Tilt automating the development environment.

## Overview

This project provides an example setup where you can debug your Go applications using Delve inside a Kubernetes cluster. The use of Tilt streamlines the process of updating and redeploying containers as you develop, allowing for real-time debugging and development.

## Prerequisites

Before you begin, ensure you have the following installed:
- Docker
- kind
- kubectl
- ctlptl
- Tilt
- Go

A Make target has been provided to bring up the environment: `make tilt`