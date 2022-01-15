# envaddr
Get Listen Address From env

``` shell
go get -u github.com/lemon-mint/envaddr
```

``` go
func Get(defaultAddr string) string
```

This function attempts to get the address from the environment variables `PORT`, `HOST` and `IP`.

If none of these are set, it returns the default address.
