/*

Package internal contains packages with restricted visibility

| Package | Description                                                                                                    |
|---------|----------------------------------------------------------------------------------------------------------------|
| assets  | provides assets for manual tests                                                                               |
| build   | contains elements used during the build process (example: utils to execute cmds, environment metadata)         |
| debug   | interfaces Go runtime debugging facilities making these facilities available through the CLI and RPC subsystem |
| jsre    | provides the execution environment for Javascript                                                              |
| kusdapi | provides an API to access Kowala related information                                                           |
| web3ext | Kowala specific web3.js extensions                                                                             |

*/

package internal
