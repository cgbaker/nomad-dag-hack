task "beta" {
  driver = "raw_exec"
  config {
    command = "sleep"
    args = ["7"]
  }
}
