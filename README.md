# vectors
Support for vectors

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