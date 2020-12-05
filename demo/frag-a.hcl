meta {
  hey = "sup"
}

task "alpha" {
  driver = "raw_exec"
  config {
    command = "sleep"
    args = ["5"]
  }
}

task "alpha-prime" {
  driver = "raw_exec"  
  config {
    command = "sleep"
    args = ["5"]
  }
}
