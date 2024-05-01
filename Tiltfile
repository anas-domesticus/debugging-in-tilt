# Here we set the build target to 'debug' stage.
docker_build(
  'example',
  'cmd/example',
  dockerfile='cmd/example/Dockerfile',
  target='debug'
)

# Apply the Kubernetes deployment located at 'k8s/deployment.yaml'.
k8s_yaml('k8s/deployment.yaml')

# Here we need a port forward from local machine to the Kubernetes resource at port 40000.
k8s_resource(
  'example-deployment',
  labels=['workload'],
  port_forwards=40000
)