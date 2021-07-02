# Traefik plugin: Vault Auth

### Prerequisites

* A valid [Traefik Pilot](https://pilot.traefik.io) token for your Traefik instance.
* A running Vault server in which one or more users are configured.

### Demo
You can found a demonstration Docker Compose file (`docker-compose.demo.yml`) in the repository root. 


```shell
TRAEFIK_PILOT_TOKEN="xxxx" docker-compose up -d --build 
```
This will launch:
* A vault (http://localhost:8200) (Vault and consul).
* A Traefik instance with [dashboard](http://traefik.localhost)
* A [`whoami` instance](http://whoami.localhost)

Once all containers are started and healthy, you can use the Vault Console to create your users (http://localhost:8200).

### Installation
Declare it in the Traefik configuration:

**YAML**
```yaml
pilot:
  token: "xxxx"
experimental:
  plugins:
    traefik-vault-auth:
      moduleName: "github.com/galaxias/traefik-vault-auth"
      version: "v0.2.3"
```

### Configuration

**YAML**
```yaml
    middlewares:
      my-traefik-vault-auth:
        plugin:
          traefik-vault-auth:
            customRealm: Use a valid Vault user to authenticate
            vault:
              routes:
                login: /v1/metadata/data/core
              token: s.XgSNXNFGcXyhs7a5Uu1gg806
              url: http://127.0.0.1:8200
```