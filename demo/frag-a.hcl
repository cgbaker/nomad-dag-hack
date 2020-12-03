task "alpha" {
  driver = "raw_exec"
  config {
    command = "sleep"
    args = ["10"]
  }
}
