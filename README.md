# Captain

## Introduction

Captain - Convert your Git workflow to Docker containers ready for Continuous Delivery

Define your workflow in the `captain.yaml` and use captain to your Continuous Delivery service to create containers for each commit, test them and push them to your registry only when tests passes.

* Use `captain build` to build your Dockerfile(s) of your repository. If your repository has local changes the containers will only be tagged as *latest*, otherwise the containers will be tagged as *latest*, *COMMIT_ID* & *BRANCH_NAME*. Now your Git commit tree is reproduced in your local docker repository.
* Use `captain test` to run your tests
* Use `captain push` to send selected images to the remote repository

From the other side, you can now pull the feature branch you want to test, or create distribution channels (such as 'alpha', 'beta', 'stable') using git tags that are propagated to container tags.

## Installation

### Github Releases

`captain` binaries can be download in each github release

### Homebrew

```
$ brew install guilhem/homebrew-tap/bump
```

### Docker

use image [`guilhem/captain`](https://hub.docker.com/repository/docker/guilhem/captain)

## Captain.yml Format

Captain will automatically configure itself with sane values without the need for any pre-configuration, so that it will work in most cases. When it doesn't, the `captain.yml` file can be used to configure it properly. This is a simple YAML file placed on the root directory of your git repository. Captain will look for it and use it to be configured.

Here is a full `captain.yml` example:

```yaml
hello-world:
  build: Dockerfile
  image: guilhem/hello-world
  pre:
    - echo "Preparing hello-world"
  post:
    - echo "Finished hello-world"
hello-world-test:
  build: Dockerfile.test
  image: guilhem/hello-world-test
  pre:
    - echo "Preparing hello-world-test"
  post:
    - echo "Finished hello-world-test"
  test:
    - docker run -e NODE_ENV=TEST harbur/hello-world-test node mochaTest
    - docker run -e NODE_ENV=TEST harbur/hello-world-test node karmaTest
hello-world-with-context:
  context: myDir
  build: Dockerfile
  image: guilhem/helly-context
hello-world-with-dependancies:
  build: Dockerfile.sup
  wants:
    - hello-world
project-with-build-args:
  build: Dockerfile
  image: harbur/buildargs
  build_arg:
    keyname: keyvalue
```

### image

The location of the Dockerfile to be compiled.

When auto-detecting, the image will be re-constructed by the following rules:
- `Dockerfile`: `username`/`parent_dir`
- `Dockerfile.*`: `username`/`parent_dir`.`parsed_suffix`

Where

- `username` is the host's username
- `parent_dir` is the Dockerfile's parent directory name
- `parsed_suffix`: is the suffix of the Dockerfile parsed with the following rules:
  - Lower-cased to avoid invalid repository names (Repository names support only lowercase letters, digits and _ - . characters are allowed)

```yaml
image: guilhem/hello-world
```

### build

The relative path of the Dockerfile to be used to compile the image. The Dockerfile's directory is also the build context that is sent to the Docker daemon.

When auto-detecting it will walk current directory and all subdirectories to locate Dockerfiles of the following format:

- `Dockerfile`
- `Dockerfile.*`

The build path will be reconstructed automatically to compile the Dockerfile. The build context will be the directory where the Dockerfile is located.

Note: If more than one Dockerfiles are needed on specific directory, suffix can be used to separate them and share the same build context.

```yaml
build: Dockerfile
build: Dockerfile.custom
build: path/to/file/Dockerfile
build: path/to/file/Dockerfile.custom
```

### test

A list of commands that are run as tests after the compilation of the specific image. If any command fail, then captain stops and reports a non-zero exit status.

```yaml
test:
  - docker run -e NODE_ENV=TEST guilhem/hello-world-test node mochaTest
  - docker run -e NODE_ENV=TEST guilhem/hello-world-test node karmaTest
```

### pre

A list of commands that are run as preparation before the compilation of the specific image. If any command fail, then captain stops and reports a non-zero exit status.

```yaml
pre:
  - echo "Preparing compilation"
```

### post

A list of commands that are run as post-execution after the compilation of the specific image. If any command fail, then captain stops and reports a non-zero exit status.

```yaml
post:
  - echo "Reporting after compilation"
```

### build_arg

A set of key values that are passed to docker build as `--build-arg` flag. For more information about build-args see [here](https://docs.docker.com/engine/reference/commandline/build/).

```yaml
build_arg:
  keyname: keyvalue
```

## CLI Commands

### build

Builds the docker image(s) of your repository

It will build the docker image(s) described on captain.yml in order they appear on file

Flags:

```
-B, --all-branches=false: Build all branches on specific commit instead of just working branch
-f, --force=false: Force build even if image is already built
-t, --tag strinf: Tag version
```

### test

Runs the tests

It will execute the commands described on test section in order they appear on file

### push

Pushes the images to remote registry

It will push the generated images to the remote registry

By default it pushes the 'latest' and the 'branch' docker tags.

Flags:

```
-B, --all-branches=false: Push all branches on specific commit instead of just working branch
-b, --branch-tags=true: Push the 'branch' docker tags
-c, --commit-tags=false: Push the 'commit' docker tags. If branch-tags=true, it also pulls the 'branch-commit' docker tags
```

### pull

Pulls the images from remote registry

It will pull the images from the remote registry

By default it pulls the 'latest' and the 'branch' docker tags.

Flags:

```
-B, --all-branches=false: Pull all branches on specific commit instead of just working branch
-b, --branch-tags=true: Pull the 'branch' docker tags
-c, --commit-tags=false: Pull the 'commit' docker tags. If branch-tags=true, it also pulls the 'branch-commit' docker tags
```

### version

Display version

Displays the version of Captain

### help

Help provides help for any command in the application.

Simply type `captain help [path to command]` for full details.

## Global CLI Flags

```
-D, --debug=false: Enable debug mode
-h, --help=false: help for captain
-N, --namespace="username": Set default image namespace
-l, --long-sha=false: Use the long git commit SHA when referencing revisions
```

## Docker Tags Lifecycle

The following is the workflow of tagging Docker images according to git state.

- If you're in non-git repository, captain will tag the built images with `latest`.
- If you're in dirty-git repository, captain will tag the built images with `latest`.
- If you're in pristine-git repository, captain will tag the built images with `latest`, `commit-id`, `branch-name`, `branch-name-commit-id`, `tag-name`. A maximum of one tag per commit id is supported.
