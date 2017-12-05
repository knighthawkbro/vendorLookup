# How to use VendorLookup
## Source code to my small little project.
VendorLookup takes in a file of MAC addresses in any format separated by newlines and spits out all the vendor specified for those mac addresses. There are options to --copy flag to read from the systems clipboard and find the vendor associated.

### Usage:
```
$ ./vendorLookup [--copy] <FileName>

Examples:
# Look up a bulk list of MAC Addresses
$ ./vendorLookup mac.txt

# Read Clipboard and output to a file
$ ./vendorLookup --copy
```

The program reads the file and fetches vendor information from [MA:CV:en:do:rs.com](https://macvendors.com/)'s API. See [https://macvendors.com/api](https://macvendors.com/api) for more information

### Enhancements to come
- ~~Test if clipboard has mac addresses saved~~ See [#78e7796](https://github.com/knighthawkbro/vendorLookup/commit/78e7796b17148cf86fd2b26fd688a5c4d32cf46d)
- Test file template before running
- ~~Add flags for file output options~~ See [#78e7796](https://github.com/knighthawkbro/vendorLookup/commit/78e7796b17148cf86fd2b26fd688a5c4d32cf46d)
- Add a verbose option instead of outputting to file
- ~~Add a license to it~~ See [#ca8345e](https://github.com/knighthawkbro/vendorLookup/commit/ca8345e944cc8df95b7b8757d4962b5307732522)
- ~~Change the name to something more descriptive~~ See [#a1596c4](https://github.com/knighthawkbro/vendorLookup/commit/a1596c4e5b21fd798901fbe38b72ab0802024822)
- Allow user to change the name of the output file with flag
- Read from STDIN or Output to STDOUT
- Save output as CSV