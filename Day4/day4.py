def main():
    # Read Input File
    #f = open("test.txt", "r")
    f = open("input_Part1.txt", "r")

    rawCards = []

    sumOfPoints = 0

    for row in f:
        row = row.replace("\n", "")
        card = row.split(": ", 2)[0].replace("Card ", "")
        cards = row.split(": ")
        numbers = row.split(": ")[1]
        winningNumbers = numbers.split(" | ")[0]
        winningNumbers = winningNumbers.split()
        havingNumbers = numbers.split(" | ")[1]
        havingNumbers = havingNumbers.split()
        rawCards.append(card)

        wins = 0
        for n in havingNumbers:
            for w in winningNumbers:
                if (n == w):
                    if wins == 0:
                        wins = 1
                    else:
                        wins = wins * 2

        sumOfPoints = sumOfPoints + wins
   
    print(sumOfPoints)





main()