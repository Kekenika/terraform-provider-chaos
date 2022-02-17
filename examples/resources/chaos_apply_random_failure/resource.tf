resource "chaos_apply_random_failure" "my_failure" {
  name         = "my-failure"
  failure_rate = 0.5
}
