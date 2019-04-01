workflow "Test" {
  on = "push"
  resolves = "Test on Travis CI"
}

action "Test on Travis CI" {
  uses = "travis-ci/actions@master"
  secrets = ["TRAVIS_TOKEN"]
}

workflow "gorelease" {
  on = "release"
  resolves = "goreleaser"
}

action "Docker Registry" {
  uses = "actions/docker/login@8cdf801b322af5f369e00d85e9cf3a7122f49108"
  secrets = ["DOCKER_USERNAME", "DOCKER_PASSWORD"]
}

action "goreleaser" {
  uses = "docker://goreleaser/goreleaser"
  secrets = [
    "GORELEASER_GITHUB_TOKEN",
  ]
  args = "release"
  needs = ["Docker Registry"]
}
