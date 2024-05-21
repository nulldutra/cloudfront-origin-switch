# cloudfront-origin-switch
A simple tool to switch CloudFront origin

## Usage

```
A simple tool to switch origins on CloudFront

Usage:
  cloudfront-origin-switch [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        Commands to list CloudFront IDs and Origins
  switch      Switch origins on CloudFront distribution

Flags:
  -h, --help     help for cloudfront-origin-switch
  -t, --toggle   Help message for toggle

Use "cloudfront-origin-switch [command] --help" for more information about a command.
```

## Examples

Listing distributions

```sh
cloudfront-origin-switch list distributions
```

Listing origins

```sh
cloudfront-origin-switch list origins --id E1HSDX49KSFRDA
```

Switching origin traffic

```sh
cloudfront-origin-switch switch --id E1B4MZC46KTIDU --from production --to fallback
```
