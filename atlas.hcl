// C:/atlas.hcl

// Define an Atlas variable to hold the database URL.
// This will be populated from an environment variable that you'll set from your .env file.
variable "db_url" {
  type = string
}

// Define your local development environment
env "local" {

  dev = "docker://postgres/15/dev_atlas"

  migration {
    dir = "file://migrations"
  }


  url = var.db_url

  dialect = "postgres"
}