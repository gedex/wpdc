wpdc
====

WordPress Deprecated Checker.

Checks deprecated functions and classes as listed in following sources:

```
// From core.
https://core.trac.wordpress.org/browser/trunk/src/wp-includes/deprecated.php?format=txt
https://core.trac.wordpress.org/browser/trunk/src/wp-admin/includes/deprecated.php?format=txt
https://core.trac.wordpress.org/browser/trunk/src/wp-includes/pluggable-deprecated.php?format=txt
https://core.trac.wordpress.org/browser/trunk/src/wp-includes/ms-deprecated.php?format=txt
https://core.trac.wordpress.org/browser/trunk/src/wp-admin/includes/ms-deprecated.php?format=txt

// WooCommerce.
https://raw.githubusercontent.com/woothemes/woocommerce/master/includes/wc-deprecated-functions.php
```

## Package

Following is the package:

```
github.com/gedex/wpdc
```

that's used by the command line `wpdc`.

## Install the Command Line Program

```
go install github.com/gedex/wpdc/cmd/wpdc
```

Example of running `wpdc`:

```
$ wpdc -h
usage: wpdc [path ...]

$ wpdc /path/to/cloned/o2
Found 4 deprecateds being used in o2/inc/fragment.php:
* line 86 uses deprecated `next_post` as listed in https://core.trac.wordpress.org/browser/trunk/src/wp-includes/deprecated.php?format=txt
* line 87 uses deprecated `next_post` as listed in https://core.trac.wordpress.org/browser/trunk/src/wp-includes/deprecated.php?format=txt
* line 89 uses deprecated `next_post` as listed in https://core.trac.wordpress.org/browser/trunk/src/wp-includes/deprecated.php?format=txt
* line 91 uses deprecated `next_post` as listed in https://core.trac.wordpress.org/browser/trunk/src/wp-includes/deprecated.php?format=txt
------------

Found 1 deprecateds being used in o2/o2.php:
* line 325 uses deprecated `get_settings` as listed in https://core.trac.wordpress.org/browser/trunk/src/wp-includes/deprecated.php?format=txt
------------
```

Args are optional path to the PHP files, if no arg is provided it consumes stdin.

## Credit

* https://github.com/stephens2424/php for the lexer.
