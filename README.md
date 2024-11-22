# vectors
Support for vectors

[![Open in Dev Containers](https://img.shields.io/static/v1?label=Dev%20Containers&message=Open&color=blue&logo=visualstudiocode)](https://vscode.dev/redirect?url=vscode://ms-vscode-remote.remote-containers/cloneInVolume?url=https://github.com/SPANDigital/vectors)
![Dev Go Action Workflow Status](https://img.shields.io/github/actions/workflow/status/spandigital/vectors/go.yml?branch=develop&label=dev)
![Main Go Action Workflow Status](https://img.shields.io/github/actions/workflow/status/spandigital/vectors/go.yml?branch=main&label=main)
![Tag](https://img.shields.io/github/v/tag/SPANDigital/vectors)

## Alias

Vector is an alias for a slice of floats
```
type Vector []float64
```

## Functions

### Cosine Similarity

```
func CosineSimilarity(a Vector, b Vector) (cosine float64, err error)
```