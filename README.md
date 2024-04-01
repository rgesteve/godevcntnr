# Testing devcontainers (for golang) using podman on Windows

1. Make sure you have Podman running on your Windows box.  A quick check is to do
```bash
podman version
```
which should report entries for "Podman client" and "Podman Engine".  If this does not work, but you had Podman working, you may have forgotten to restart the WSL instance that actually runs the Podman engine.  You can do this by running
```bash
podman machine start
```

1. Create the devcontainer.  If you don't have the CLI installed, you can do:
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
