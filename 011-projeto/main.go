package main

import "fmt"

func calcularDigito(cpf string, digito int) int {
	soma := 0
	for i := 0; i < len(cpf); i++ {
		num := int(cpf[i] - '0') // transforma char em número
		soma += num * (digito - i)
	}
	resto := (soma * 10) % 11
	if resto == 10 {
		return 0
	}
	return resto
}

// Recebe o CPF no formato 00000000000 e se verdadeiro retorna como 000.000.000-00
func IsCPF(cpf string) (string, bool) {
	if len(cpf) != 11 {
		return "", false
	}
	// Calcular digitos verificadores
	d1 := calcularDigito(cpf[:9], 10)
	d2 := calcularDigito(cpf[:9] + fmt.Sprint(d1), 11)


	if d1 != int(cpf[9]-'0') || d2 != int(cpf[10]-'0') {
		return "", false
	}

	cpfFormatado := fmt.Sprintf("%s.%s.%s-%s", cpf[0:3], cpf[3:6], cpf[6:9], cpf[9:11])

	return cpfFormatado, true
}


func main() {
	cpf := "50021864837"
	cpfFromatado, verif := IsCPF(cpf)
	if verif == false {
		fmt.Println("O CPF não é váido")
		return
	}
	fmt.Printf("CPF formatado: %s \n", cpfFromatado)
}