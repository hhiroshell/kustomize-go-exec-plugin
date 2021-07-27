Kustomize Go Exec Plugin
===
Sample kustomize exec plugins written in Go.

Kustomize supports two types of plugins (Go / Exec), but the Go plugin is "[very annoying](https://github.com/kubernetes-sigs/kustomize/issues/3574)" because it is realized using the Go plugin system.
This repository is trying to write some Kustomize exec plugins in Go.

This repository contains the following simple plugins:

- Secret Generator:
    - Generates a manifest of kubernetes secret. You can specify the secret name, namespace and literals(key-value pairs)
      of the data field.
- Line Insertion Transformer:
    - Inserts lines into base manifests. This plugin inserts specified strings above lines those include the anchor text. 

Usage
---

### Build and Run
First, you have to build Go binaries those will be used as Kustomize exec plugin.
You can do it only by executing `make` job at the repository root dir.

```console
$ git clone git@github.com:hhiroshell/kustomize-go-exec-plugin.git && cd kustomize-go-exec-plugin

$ make
```

Then, set the environment variable `KUSTOMIZE_PLUGIN_HOME` to the path of the `plugin` directory.

```console
$ export KUSTOMIZE_PLUGIN_HOME=$(pwd)/plugin
```

Now, you can use the "SecretGenerator" sample plugin by executing the `kustomize` command as follows.

```console
$ kustomize build --enable-alpha-plugins example/secret-generator
```

### How to write the kustomization.yaml
Sample plugin binaires can run as Go CLI commands. You can see all flags of plugins by executing `help` subcommands.

```console
$ plugin/hhiroshell.github.com/v1/secretgenerator/SecretGenerator --help
Kustomize exec plugin that generates a manifest of kubernetes secret written in yaml format.
You can specify the secret name, namespace and literals(key-value pairs) of the data field, via CLI flags.

Usage:
  SecretGenerator [flags]

Flags:
  -h, --help                  help for SecretGenerator
      --literal stringArray   Literal key-value pairs used as data of the Secret. (e.g. --literal key=value)
      --name string           Name of the generated Secret
      --namespace string      Namespace of the generated Secret (default "default")
```

All options are available by writing them in the `argsOneLiner` field of the plugin’s configuration file referenced from kustomization.yaml.

- plugin's configuration file (secret-generator.yaml)

```yaml
apiVersion: hhiroshell.github.com/v1
kind: SecretGenerator
metadata:
  name: secret-generator
argsOneLiner: |
  --name exmaple-name
  --namespace example-namespace
  --literal hoge=fuga
  --literal foo=bar
```

- kustomization.yaml

```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

generators:
- secret-generator.yaml
```

For more detail about plugin configurations, see the official [Kustomize plugins guide](https://kubectl.docs.kubernetes.io/guides/extending_kustomize/).
