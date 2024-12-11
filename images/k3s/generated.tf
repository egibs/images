# DO NOT EDIT - this file is autogenerated by tfgen

variable "test_repository" {
  description = "The docker repo root to use for sourcing test images."
}

output "summary" {
  value = {
    "tags" = merge(module.tagger.imagetags, module.tagger-static.imagetags)
  }
}

