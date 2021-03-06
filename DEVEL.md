# IO-Gateway development process

If you are one of those heroic people wanting to help and you can deal with bugs, crashes and the occasional complete destruction of you computer, you can follow with those instructions.

After you read this document, do not forget to [check also this other document](DEBUG.md) for debugging tips.

## Prepare the Continuous Integration on DockerHub and GitHub

In order to run properly on GitHub and build images on DockerHub you need

- an account on DockerHub with a proper access token
- an account in GitHub configured to build in DockerHub

Follow this procedure to setup Docker Hub

- Register on `hub.docker.com` an account. 
- Go on yor account then security 
- generate a new access token and take note of it

Now, the GitHub part:

- Fork `io-gateway` the repository on GitHub
- Add to the GitHub secrets `DOCKER_USER`  and `DOCKER_TOKEN`
- Set the default branch `develop`
- enable actions for it (by default for a branch they are not)

## Prerequisites for your local system

Supported development platforms are:
- Linux Ubuntu 18.x
- Mac OS Catalina
- Windows 10 build 2003 with WSL2 and Ubuntu 18.04

Also your workstation needs at least 8gb of memory.

It may work on other configurations but it is not tested.

Before doing anything, you have to setup your environment as follows

### Linux Ubuntu 18.x

You need to install docker [as described here](https://docs.docker.com/engine/install/ubuntu/)

### Mac OSX Catalina

You need to:
- Install brew [as described here](https://brew.sh/)
- Install command line tools with the command:  `xcode-select --install`
- Install [Docker for Mac](https://docs.docker.com/docker-for-mac/install/) and give at leas 4GB of memory

### Windows 10

Windows configuration is a bit more complex. The steps are:
- update Windows 10 to build 2009
- install Ubuntu-18.04 under WSL2
- install Docker for Windows
- use docker from WSL2-Ubuntu
- create an ssh proxy to access localhost

More details below.

#### Update to build 2009

You need to update Windows 10 at lest to build 2009 in order to have WSL2.

Check here to read more informations [to update windows are here](https://support.microsoft.com/en-us/help/4027667/windows-10-update)

#### Install Ubuntu-18.04 under WSL2

Go into the Microsof Store and seach for Ubuntu. Install version 18.04 and launch it.

By default it may be installed on WSL1, so you need to ensure it works with WSL2.

Open PowerShell, ensure you see `Ubuntu-18.04` with the command `wsl -l` then type: `wsl --set-version Ubuntu-18.04 2`.

Now you can enter in the distro with the command `wsl -d Ubuntu-18.04`

#### Install Docker Desktop for Windows

Instructions to install Docker Desktop for Windows [are here](https://docs.docker.com/docker-for-windows/install/). You need at least version 2.3.0.2

#### Use docker for windows as docker backend in WSL 2

This step is critical.

After installation follow [these instructions](https://docs.docker.com/docker-for-windows/wsl/) to use the Docker running in Windows as the Docker to use in WSL.

## Setup the local development environmewnt

If the prerequisites are satisfied, you can setup and test the development environment as follows:

```
git clone https://github.com/pagopa/io-gateway
cd io-gateway
bash setup.sh <DOCKER_USER> <DOCKER_TOKEN>
source source-me-first
make
```

If all the test passes, congratulations, you development environment is ready.


## Daily development tasks

In general you do not need to repeat the entire setup procedure every day.

If everything was setup, just type:  `source source-me-first` to setup the environment.

You can update images with:

```
make build
```

You can start development services. Do:

```
./iogw/iogw init --wskprops
./iogw/iogw start --skip-pull-images --skip-ide
```

Note the `--wskprops` that will setup locally the configuration file to access and deploy actions in OpenWhisk.
Note also the `--skip-pull-images` that will use the local copy without downloading images from the net

If you are lucky, your browser will open and you will get the io-gateway user interface.

Try also `wsk action list` and you should see the list of actions

## Admin User Interface

To develop the admin user interface, do:

```
cd admin
make devel
```

Now you have a development version of the UI in `http://localhost:5000`.

**NOTE:** on windows with WSL2 you will need to setup a tunnel to reach that port. See the section on installing on Windows.

Code is writtend javascript, based on [Svelte](https://svelte.dev/). Sources are under `src`

You can edit code and it will automatically reload.

Do `make build` from the top level to embed the code you modified in the image.

## Backend Actions

The server side part of the backend are [OpenWhisk](http://openwhisk.apache.org), action written in various languages.

Source code is under `admin/actions/src`

You can edit sources.

You can deploy actions after changes with `make deploy`

## The CLI

The CLI, that is used to launch and initialize the environment is under `iogw` and it is written in the Go programming language.

To develop the cli, do `cd iogw` and use `make`.

The cli embeds the version number of the images to refer to them correctly.

In development, the version is just the branch name.

When you push to master, the CI builds images and pushes to the Docker Hub, so the CLI can retrieve the version stored in Docker Hub.

There is a test suite for the CLI and you generally should ensure tests are passed with `make test`.

## Snapshot and Releases

Images are pushed to docker hub only by the CI, and they uses a secret token, so to get your changes in the public images on docker hub you need to open a Pul Requests.

If you have the token (reserved to release manager) you can do, after a `docker login` also:

- `make push` to push images tagged with current branch
- `make snapshot` to create a "snapshot" release (in format `%Y.%m%d.%H%M-snapshot`)
- `git tag <release> ; git push --tags` to create a release (a manual process)

## Base images

The server, the ide and the installer all uses some base images, that are pretty stable and change rarely.

If you really need to change them, you need to get a password to write in some public repository (`pagopa` on DockerHub currently)

`cd images ; make`

They are tagged with the date you built them. If you change the tag you will also have to change the tag in `admin/Dockerfile` and `ide/Dockerfile`

The process is manual because it is infrequent and needs awareness of what you are doing.
