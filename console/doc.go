/*
Package console implements an interactive javascript running environment

The current implementation of the console relies on the Ethereum Javascript framework (web3) for apis that are covered in it
and also on extensions to the same framework available at internal/web3ext.
Note that any update to the backend must involve also some updates to either web3 or the extensions.

Interactive Use - Console

Once you have the kusd client installed, using the console is merely a case of executing the "kusd <your_config> console"
command in a terminal - the console displays the available modules upon start.

Non-Interactive Use

It's also possible to execute files in the Javascript interpreter. The "console" and "attach" subcommand accept the "--exec" argument which
is a javascript statement.

Limitations

Kowala JSRE uses the Otto VM which has some limitations:

* "use strict" will parse, but does nothing.
* The regular expression engine (re2/regexp) is not fully compatible with the ECMA5 specification.
*/

package console
