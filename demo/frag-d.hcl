task "delta" {
  driver = "raw_exec"
  config {
    command = "sleep"
    args = ["3"]
  }
}
