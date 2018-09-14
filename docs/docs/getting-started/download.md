# Downloads

## Binaries

Binary distributions are available for Linux, Mac OS X and Windows (though we
don't currently offer support for the Windows build). You'll need a binary, 
as opposed to a Docker image, in order to use USB hardware wallets (including
Ledger devices).

After downloading a binary release suitable for your system, please follow the
[Connecting to the Kowala Network](/getting-started/testnet/#connecting-to-the-kowala-network)
instructions.

### Supported, stable versions

 | Operating system           | File                                                                                                                              |   
 | -------------------------- | --------------------------------------------------------------------------------------------------------------------------------- | 
 | Mac OS X                   | [kcoin-stable-osx-10.6-amd64.zip](https://s3.amazonaws.com/releases.kowala.tech/kcoin-stable-osx-10.6-amd64.zip)                  |
 | Linux 64bits               | [kcoin-stable-linux-amd64.zip](https://s3.amazonaws.com/releases.kowala.tech/kcoin-stable-linux-amd64.zip)                        |
 | Linux ARM 64bits           | [kcoin-stable-linux-arm64.zip](https://s3.amazonaws.com/releases.kowala.tech/kcoin-stable-linux-arm64.zip)                        |

### Unsupported, stable versions

 | Operating system           | File                                                                                                                              |   
 | -------------------------- | --------------------------------------------------------------------------------------------------------------------------------- | 
 | Windows                    | [kcoin-stable-windows-4.0-amd64.exe.zip](https://s3.amazonaws.com/releases.kowala.tech/kcoin-stable-windows-4.0-amd64.exe.zip)    |

## Docker image

The easiest way to run the mining client is using Docker. The latest docker
image is always available via
[Dockerhub](https://hub.docker.com/r/kowalatech/kusd/). You can get the kUSD
version by running 

``` docker pull kowalatech/kusd ```

in a terminal.

### Dockerhub tags

There are two main tags:

| Tag name  |                      Purpose                      |
|-----------|---------------------------------------------------|
| `:latest` | The most up to date, stable version of the client |
| `:dev`    | Cutting-edge development version of client.       |

There are also granular tags for each verision. See the [full
list](https://hub.docker.com/r/kowalatech/kusd/tags/).

