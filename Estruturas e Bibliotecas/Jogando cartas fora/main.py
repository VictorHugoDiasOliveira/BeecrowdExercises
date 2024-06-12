# Dada uma pilha de n cartas enumeradas de 1 até n com a carta 1 no topo e a carta n na base.  A seguinte operação é ralizada enquanto tiver 2 ou mais cartas na pilha.

# Jogue fora a carta do topo e mova a próxima carta (a que ficou no topo) para a base da pilha.

# Sua tarefa é encontrar a sequência de cartas descartadas e a última carta remanescente.

# Cada linha de entrada (com exceção da última) contém um número n ≤ 50. A última linha contém 0 e não deve ser processada. Cada número de entrada produz duas linhas de saída. A primeira linha apresenta a sequência de cartas descartadas e a segunda linha apresenta a carta remanescente.

# Entrada
# A entrada consiste em um número indeterminado de linhas contendo cada uma um valor de 1 até 50. A última linha contém o valor 0.

# Saída
# Para cada caso de teste, imprima duas linhas. A primeira linha apresenta a sequência de cartas descartadas, cada uma delas separadas por uma vírgula e um espaço. A segunda linha apresenta o número da carta que restou. Nenhuma linha tem espaços extras no início ou no final. Veja exemplo para conferir o formato esperado.

class playing_cards():
    def generate(self, input: int) -> list:
        self.discarded = []
        self.cards = []
        for i in range(input):
            self.cards.insert(0, i+1)
        return self.cards

    def discard(self) -> list:
        self.discarded.append(self.cards[len(self.cards)-1])
        self.cards.pop()
        self.cards.insert(0, self.cards[len(self.cards)-1])
        self.cards.pop()
        if len(self.cards) == 1:
            return self.cards, self.discarded
        return self.discard()
    
    def show_info(self) -> None:
        print(f"Discarded cards: {', '.join(str(x) for x in self.discarded)}")
        print(f"Remaining card: {self.cards[0]}")

play = playing_cards()

while True:
    number_of_cards = int(input())
    if number_of_cards == 0:
        break
    else:
        play.generate(number_of_cards)
        play.discard()
        play.show_info()