provider "null" {
  # No specific configuration needed for null provider
}

resource "null_resource" "podman_cleanup" {
  provisioner "local-exec" {
    command = <<EOT
      # Run podman-compose down
      podman-compose down

      # Get podman images
      IMAGES=$(podman images --format "{{.ID}} {{.Repository}}" | grep "site-portfolio_web" | awk '{print $1}')

      # Remove the images if found
      if [ ! -z "$IMAGES" ]; then
        podman rmi $IMAGES
      else
        echo "No site-portfolio_web image found."
      fi
    EOT
  }
}