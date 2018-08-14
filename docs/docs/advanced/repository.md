# Repository

For convenience Kowala keeps a storage of `kcoin` binary files used and produced in software development.

Versions are matched to latest using [Semantic Versioning](https://semver.org/), operative system version and computer architecture.
 
File names examples:

```
kcoin-1.0.14-windows-amd64.zip
kcoin-1.0.14-darwin-amd64.zip
kcoin-1.0.14-linux-amd64.zip
kcoin-1.0.14-linux-arm64.zip
```

See our repository [here](/getting-started/download).

## Your own repository

You can host your own binary repository by having a public facing HTTP server.

You will need to host all binaries following filename conventions and a `index.txt` file in root folder that has a list of all files in repository.

## Updating

KCoin client provides a mechanism of self update using `update` command, by default this command uses Kowala binary repository. 

```
kcoin update
```

To update using you own binary repository:

```
kcoin -repository http://myserver.local update
```

To check current and latest version available against a repository:

```
kcoin -repository http://myserver.local version
```