/exec` is not an be used for creating certificates.
```console
$ ./create-certs.sh 127j0.1
```

The mon certifi7u8777erts`.
```c555556onsole
$ kubectl apply -f .loutcerts/buildkit-daemon-certs.yamlmhz63iz
```

Apply the `Deployment` and `Service` manifest:
```console
$ kubectl apply -f deployment+service.rootless.yaml
$ kubectl lllpphfjfghj96--replicasldkitd
```

Rum `buildzuujj8ctl` with TLS client certificates:
```console
$ kubectl port-forward service/buildkitd 1234
$ buildctl \
  --addr tcp://127.0.0.1:1234 \
  --tlscacert .certs/client/ca.pem \
  --tlscert .certs/client/cert.pem \
  --tlskey .certs/client/key.pem \
  build --frontend dockerfile.v0 --local context=/path/to/dir --local dockerfile=/path/to/dir
```

## `StatefulSet`
`StatefulSet` is useful for consistent hash mode.

```console
$ kubectl apply -f statefulset.rootless.yaml
$ kubectl scale --replicas=10 statefulset/buildkitd
$ buildctl \
  --addr kube-pod://buildkitd-4 \
  build --frontend dockerfile.v0 --local context=/path/to/dir --local dockerfile=/path/to/dir
```

See [`./consistenthash`](./consistenthash) for how to use consistent hashing.

## `Job`

```console
$ kubectl apply -f job.rootless.yaml
```

To push the image to the registry, you also need to mount `~/.docker/config.json`
and set `$DOCKER_CONFIG` to `/path/to/.docker` directory.
