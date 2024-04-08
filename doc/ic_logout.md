## ic logout

Log out

```
ic logout [flags]
```

### Options

```
      --oidc-issuer-url string                       Issuer URL for the OIDC Provider (default "http://localhost:8080/realms/test")
      --oidc-client-id string                        OIDC Client ID (default "inventory-cli")
      --oidc-grant-type string                       OIDC Authorization grant type. One of (authcode-browser|authcode-keyboard) (default "authcode-browser")
      --oidc-redirect-url-hostname string            [authcode-browser] Hostname of the redirect URL (default "localhost")
      --oidc-redirect-url-authcode-keyboard string   [authcode-keyboard] Redirect URL when using authcode keyboard (default "urn:ietf:wg:oauth:2.0:oob")
      --oidc-auth-bind-addr string                   [authcode-browser] Bind address and port for local server used for OIDC redirect (default "localhost:18000")
  -h, --help                                         help for logout
```

### Options inherited from parent commands

```
  -i, --interactive string   Run in interactive mode. One of (yes|no|auto) (default "auto")
  -l, --log-level string     Set log level (default "info")
  -s, --server string        URL for the inventory server. (default "http://localhost:8086")
```

### SEE ALSO

* [ic](ic.md)	 - Inventory Client

###### Auto generated by spf13/cobra on 8-Apr-2024