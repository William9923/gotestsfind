# gotestsfind

Go test finder. A program inspired heavily from `gotestsum`.

WIP Feature:
- Initial always do list for the cache for whatever command... "-list"

- Checkout the json output result, to remove the grep and move to jq instead if possible

- Move to flag based system.

Flag that is planned:
- `-d` : debug the subtest via delve
- `-v` : verbose (for error). Can also be used to provide more details on the test...
- `-l` : list. Print all currently cache test
- `-c` : clear cache
- `-r` : refresh cache
- `-h` : help command

- Custom go test flag: https://github.com/gotestyourself/gotestsum?tab=readme-ov-file#custom-go-test-command

- More verbose logging...
