<h1 align="center">Trench Coat Movie Star: Proof of Concept Party Game</h1>

THIS REPOSITORY IS A PROOF OF CONCEPT FORKED FROM ANOTHER PROJECT.
No new major functionalities have been added, as this repository is only to showcase the possibilities for a new project. 
To access the original forked project, "OTTOMH", navigate to: https://github.com/cis3296f22/ottomh

## Local installation

If you prefer to play the game locally instead of on the live server at https://ottomh.herokuapp.com/, then you'll need to first install the app by either cloning the main branch of this repo or download the latest released zip files under "releases".

### Installing

To download the game for local use, clone the repo or download the source from the 'releases' page, then install dependencies using.

To clone the repo:
```bash
git clone https://github.com/cis3296f22/ottomh.git
```

### Building

First install the dependencies: 
1. Install NPM packages with `npm install`.
2. Install Go packages with `go get`.

To run the project:


1. Bundle JS code with `npm run build`.
2. Run the server with `go run server.go`.
3. Go to [localhost:8080](http://localhost:8080/)

By default, the server listens on port 8080. If you would like to use another port, set the $PORT environment variable.

## Contribution

All contributions are appreciated and help support the project!

If you would like to contribute to this project, please fork this repo, then make a pull request explaining your changes and contribution. Before making a pull request, please make sure that all tests pass. Testing instructions are under the [Running tests](#running-tests) sections. 

OTTOMH is built using ReactJs and Go, so you will need the `npm` and `go` commands in order to build and run the project as well as installing the necessary dependencies. Please see the build instructions under the '[Building](#building)' section of this README.

### Issues and feature requests

To report any issues and/or bugs, please open an issue and include the following information so we can better understand and solve the problem:
* State what the issue is
* What did you do to encounter this issue so that we can replicate it
* What browser you were using and what version was it
* Additionally, if you're running the game on a local machine, what operating system and version were you using

If you would like to make a feature request, please include in the title "Feature request:". For example, "Feature request: add light mode".

### Running Tests

[tests.sh](tests.sh) tests the backend with coverage and tests the frontend with coverage using a variety of third-party tools. To run these tests locally you will need to take a few extra steps:
1. Install [go-test-report](https://github.com/vakenbolt/go-test-report) with `go install github.com/vakenbolt/go-test-report@latest`. Make sure to add your GOPATH to the PATH (usually the GOPATH is `~/go/bin`)
2. [Line 3](tests.sh) of tests.sh is configured to open an HTML file in the Brave Browser using Mac's `open` command. You will have to edit this line if you are on a platform other than Mac.
3. The `time_picked` at [src/components/LobbyPage/LobbyPage.js](src/components/LobbyPage/LobbyPage.js) should be set to a very small value, like `"00:03"`.

### Generating Documentation

To generate documentation for the backend, start by installing godoc by running `go install golang.org/x/tools/cmd/godoc@latest`. Then go to the application root, and run `godoc -http=:6060`. You can then visit the go docs at localhost:6060. Here, search for the section labeled "ottomh".

To generate documentation for the frontend, use the jsdoc tool. Use `npm i jsdoc --save-dev` to install jsdoc. To generate documentation for the frontend, use `jsdoc -r src`. Then you can view the documentation by opening `out/index.html`.

## Credits
* [Brain image](src/images/logo.png) by [yo_han](https://www.freepik.com/free-vector/abstract-brain-background-design_1016468.htm#query=brain&position=17&from_view=search&track=sph%22%3E) on Freepik
