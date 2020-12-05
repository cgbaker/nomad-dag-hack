task "gamma" {
  driver = "raw_exec"
  config {
    command = "sleep"
    args = ["5"]
  }
}
