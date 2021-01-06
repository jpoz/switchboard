# switchboard

Dev proxy

# how it works:

## Step 1: Make a config:

```yaml
localProxies:
    postgres:
        proxyAddress: '0.0.0.0:6543'
        connectingAddress: '127.0.0.1:5432'
    google:
        proxyAddress: '0.0.0.0:4000'
        connectingAddress: 'google.com:80'
    redis:
        proxyAddress: '0.0.0.0:6380'
        connectingAddress: '127.0.0.1:6379'
```

## Step 2: run `switchboard` with config

```shell
$ switchboard config.yml
INFO[0000] Starting conductor 🚊
INFO[0000] Listening on 0.0.0.0:6543 proxing to 127.0.0.1:5432
INFO[0000] Listening on 0.0.0.0:6380 proxing to 127.0.0.1:6379
INFO[0000] Listening on 0.0.0.0:4000 proxing to 172.217.3.206:80
```

## Step 3: use the proxy

```shell
$ psql -h localhost --port=6543
```

