# Description: List of variables which can be provided ar runtime to override the specified defaults 

variable "project_id" {
  description = "GCP Project ID"
  type        = string
  nullable    = false
}

variable "name" {
  description = "Base name to derive everythign else from"
  default     = "vulnasync"
  type        = string
  nullable    = false
}

variable "location" {
  description = "Deployment location"
  default     = "us-west1"
  type        = string
  nullable    = false
}

variable "git_repo" {
  description = "GitHub Repo"
  default     = "mchmarny/vulnasync"
  type        = string
  nullable    = false
}
