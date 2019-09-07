# Given a list of relationships check if there is a violation return true, not return false
- The function given a list of relationships with structure `var relationships [][2]uint`
- Overview of the problem of finding loops in directed graphs

# Example 
- Check list of relationships: 
  - A < B
  - B < C
  - E < B
  - C < A

```
func main() {
  var relationships [][2]uint

  relationships = append(relationships, [2]uint{'A', 'B'})
  relationships = append(relationships, [2]uint{'B', 'C'})
  relationships = append(relationships, [2]uint{'E', 'B'})
  relationships = append(relationships, [2]uint{'C', 'A'})

  result := isViolation(relationships)

  fmt.Println("List has violation", result)
}
```

Output 
```
List has violation true
```
