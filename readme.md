# GitHub Username Finder

This command-line application allows you to find vacant GitHub account usernames by generating possible usernames based on a set of options and checking their availability on GitHub.

## Usage

To use the GitHub Username Finder, simply run the following command:

```bash
go run main.go --help
```

This will display the available options and commands for the application.

__i.e.__: I found my new login (https://github.com/ic-n) by running `go run main.go -i=30 -c=2 -hypen=50 -opts="ic,n"`


### Options

- `--opts` (default: "nice,cool,frog"): Comma-separated list of possible parts in the username.
- `--iterations`, `-i` (default: 10): Number of generations to test.
- `--concats`, `-c` (default: 3): Number of concatenation operations to create the username.
- `--hyphen` (default: 100): Percentage chance of hyphen insertion.

## Example Usage

Here's an example of how to use the GitHub Username Finder:

```bash
# Generate and check 10 usernames with default options
go run main.go

# Generate and check 20 usernames with custom options
go run main.go --iterations 20 --concats 4 --hyphen 50 --opts "nice,cool,frog"
```

## Command

- `help, h`: Shows a list of commands or help for one command.

## License

This GitHub Username Finder is open-source and available under the [MIT License](LICENSE).

## Author

This CLI application was developed by Nikolai Kiselev.
