# Automatic client updates

_If you just want to download the newest version, see the [downloads
page](/getting-started/download)._

For convenience Kowala keeps a repository of `kcoin` binaries (the mining
client) which is updated as software development progresses. You can use this
repository to keep your client up to date automatically, but you don't have to:
you're free to upgrade your client manually, or even host your own repository
and use that for automatic updates. 

We just use this system to deploy mining clients on our own testnets, and also to
update our own personal mining machines, so we're pretty sure it works well.
Since all kcoin networks are decentralized and owned by their users, we can't
(and wouldn't) mandate the use of automatic updates &mdash; like everything
else we produce, you're more or less free to do whatever you like. That said,
we do recommend that you use automatic updates for the health of the network.

## Basic usage via the CLI

The kcoin client provides a mechanism of self update using `kcoin update`
command. By default this command uses Kowala binary repository, so if you want
to use that you don't need any arguments. You can just run:

``` $ kcoin update ```

To update using you another binary repository you can use the `-repository` (or
`--repository`) flag:

``` $ kcoin -repository http://myserver.local update ```

To check current and latest version available against a repository:

``` $ kcoin -repository http://myserver.local version ```

## Under the hood

Versions are maintained using [Semantic Versioning](https://semver.org/),
operating system version and computer architecture, and these are reflected in
the file name. For example:
 
``` kcoin-1.0.14-windows-amd64.zip kcoin-1.0.14-darwin-amd64.zip
kcoin-1.0.14-linux-amd64.zip kcoin-1.0.14-linux-arm64.zip ```

When you run a `kcoin update` command against any repository, the client checks
the list of files to see if a newer version is available. If it is, then it's
automatically downloaded.

_Note that major versions of the binary are **not** automatically downloaded.
We expect to release these infrequently, and they would constitute a major
change to the communications protocols or other external interfaces. As a
result, any major vrsion updates must be performed manually._

## Hosting your own repository

You can host your own binary repository using any public-facing HTTP server.

You will need to host all binaries, following the filename conventions above,
and include an `index.txt` file in root folder that has a list of all files in
repository. That's pretty much all there is to it!
