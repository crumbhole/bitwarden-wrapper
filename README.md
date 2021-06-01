# Bitwarden wrapper

A simple wrapper around a few bitwarden CLI commands for golang. Really only aimed around reading items from folders. Could be extended for more functions, but that's all I need.

Requires that the [bitwarden CLI](https://bitwarden.com/help/article/cli/) (bw) is installed and on the path and you're already authenticated - basically you'll need a valid token in the BW_SESSION environment variable.

## Calls

The only actual call you're expected to make is
```go
item, err := GetItemFromFolder(item, folder)
```
which takes items and folders by name and expects them to match those in bitwarden by name exactly.

You can just call GetItem and GetFolder if you're happy to do something with the results

## Why not do...

Basically because this is security related and reimplementation of anything is riskier than just using an existing tool for the purpose for which I needed it.

