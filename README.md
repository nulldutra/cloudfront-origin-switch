# cloudfront-switch-origin
A simple tool to switch CloudFront origin

## Usage

```
A simple tool to switch origins on CloudFront

Usage:
  cloudfront-switch-origin [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        Commands to list CloudFront IDs and Origins
  switch      Switch origins on CloudFront distribution

Flags:
  -h, --help     help for cloudfront-switch-origin
  -t, --toggle   Help message for toggle

Use "cloudfront-switch-origin [command] --help" for more information about a command.
```

## Examples

Listing distributions

```sh
cloudfront-switch-origin list distributions
```

Listing origins

```sh
cloudfront-switch-origin list origins --id E1HSDX49KSFRDA
```

Switching origin traffic

```sh
cloudfront-switch-origin switch --id E1B4MZC46KTIDU --from production --to fallback
```
