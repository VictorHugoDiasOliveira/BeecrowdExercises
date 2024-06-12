# Considerando a entrada de valores inteiros não negativos, ordene estes valores segundo o seguinte critério:

# Primeiro os Pares
# Depois os Ímpares
# Sendo que deverão ser apresentados os pares em ordem crescente e depois os ímpares em ordem decrescente.

# Entrada
# A primeira linha de entrada contém um único inteiro positivo N (1 < N <= 105) Este é o número de linhas de entrada que vem logo a seguir. As próximas N linhas conterão, cada uma delas, um valor inteiro não negativo.

# Saída
# Apresente todos os valores lidos na entrada segundo a ordem apresentada acima. Cada número deve ser impresso em uma linha, conforme exemplo abaixo.

def generate_and_separete_array() -> list:
    odd_array = []
    even_array = []

    n = int(input())
    for _ in range(n):
        number = int(input())
        if number % 2 == 0:
            odd_array.append(number)
        else:
            even_array.append(number)
    return odd_array, even_array

def sort_and_show_array(array: list, logic: bool) -> None:
    array.sort(reverse=logic)
    for numbers in array:
        print(numbers)

odd_array, even_array = generate_and_separete_array()

sort_and_show_array(odd_array, False)
sort_and_show_array(even_array, True)