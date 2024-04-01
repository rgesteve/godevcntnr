# Testing devcontainers (for golang) using podman on Windows

1. Make sure you have Podman running on your Windows box.  A quick check is to do
```bash
podman version
```
which should report entries for "Podman client" and "Podman Engine".  If this does not work, but you had Podman working, you may have forgotten to restart the WSL instance that actually runs the Podman engine.  You can do this by running
```bash
podman machine start
```

2. Create the devcontainer.  If you don't have the CLI installed, you can do:
```bash
npm install @devcontainer/cli -g
```
Then, you can create the devcontainer by running
```bash
devcontainer up --docker-path podman --workspace-path .
```

Note that the `devcontainer` CLI requires `node` to run.  If you don't have it installed, you can do it either from Visual Studio or from winget:
```bash
winget install -e --id OpenJS.NodeJS.LTS
```

(if you've installed `devcontainer` not as a global tool, but under the root of this repository, you may need to prepend `npx` to the command above)

3. Log on to the container and build the project:
```bash
npx devcontainer exec --docker-path podman --workspace-folder . bash
```
which will open an interactive shell inside the container.  You can then build the project by running
```bash
pushd /workspaces/first/src/main
go build
```
Alternatively, you can issue the build command directly from the host by running
```bash
npx devcontainer exec --docker-path podman --workspace-folder . go build -o <NAMEFOREXECUTABLE> /workspaces/first/src/main/main.go
```