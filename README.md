# How to use VendorLookup
## Source code to my small little project.
This program takes a single file name as an input and spits out a file with all the vendors associated to the mac address.

Usage
```
$ ./vendorLookup filename.txt
```

The program reads the file and fetches vendor information from [MA:CV:en:do:rs.com](https://macvendors.com/)'s API. See [https://macvendors.com/api](https://macvendors.com/api) for more information

### Enhancements to come
- Have a singe mac address as input
- Test if clipboard has mac addresses saved
- Test file template before running
- Add flags for file output options
- Add a verbose option instead of outputting to file
- ~~Add a license to it~~ See #ca8345e
- ~~Change the name to something more descriptive~~ See #TBD