---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "chaos_plan_random_failure Resource - terraform-provider-chaos"
subcategory: ""
description: |-
  Simple resource that fails on plan
---

# chaos_plan_random_failure (Resource)

Simple resource that fails on plan

## Example Usage

```terraform
resource "chaos_plan_random_failure" "my_failure" {
  name         = "my-failure"
  failure_rate = 0.5
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) Name of the resource

### Optional

- **id** (String) The ID of this resource.

