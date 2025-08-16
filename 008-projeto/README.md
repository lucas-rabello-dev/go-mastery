### Inversor de String


Código:
```go
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j:= 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
```

<table>
<tr>
    <th></th>
    <th>0</th>
    <th>1</th>
    <th>2</th>
    <th>3</th>
    <th>4</th>
</tr>
<tr>
    <td></td>
    <td>L</td>
    <td>U</td>
    <td>C</td>
    <td>A</td>
    <td>S</td>
</tr>
<tr>
    <td>Váriaveis loop 1</td>
    <td>i</td>
    <td></td>
    <td></td>
    <td></td>
    <td>j</td>
</tr>
<tr>
    <td>Váriaveis loop 2</td>
    <td></td>
    <td>i</td>
    <td></td>
    <td>j</td>
    <td></td>
</tr>
<tr>
    <td>Váriaveis loop 3</td>
    <td></td>
    <td></td>
    <td>i j</td>
    <td></td>
    <td></td>
</tr>
<tr>
    <td>Váriaveis loop 4</td>
    <td></td>
    <td>j</td>
    <td></td>
    <td>i</td>
    <td></td>
</tr>
</table>

Fim da condição do loop no quarto loop `i < j == False` agora.

E para cada iteração no loop os valores de `i` e `j` eram trocados
```go
runes[i], runes[j] = runes[j], runes[i]
```