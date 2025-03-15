# builder.tf

provider "local" {
  # Local provider to run commands on the local machine
}

resource "null_resource" "podman_compose" {
  # Run the podman-compose build and up commands locally
  provisioner "local-exec" {
    command = "podman-compose -f ${path.module}/docker-compose.yml build && podman-compose -f ${path.module}/docker-compose.yml up"
  }

  triggers = {
    # Trigger the build and up commands on any change to the docker-compose.yml file
    docker_compose_file = "${path.module}/docker-compose.yml"
  }
}
