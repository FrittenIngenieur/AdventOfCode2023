def mainOne():
    # Read Input File
    #f = open("test.txt", "r")
    f = open("input_Part1.txt", "r")

    rawCards = []

    sumOfPoints = 0

    for row in f:
        row = row.replace("\n", "")
        card = row.split(": ", 2)[0].replace("Card ", "")
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

def mainTwo():
    # Read Input File
    #f = open("test.txt", "r")
    f = open("input_Part1.txt", "r")
    g = open("input_Part1.txt", "r")

    rawCards = []

    sumOfPoints = 0

    rowCounter = 1
    amountCards = []
    for x in g:
        counter = [rowCounter, 1]
        amountCards.append(counter)
        rowCounter = rowCounter + 1
    
    numberOfCards = rowCounter

    rowCounter = 1

    for row in f:
        row = row.replace("\n", "")
        card = row.split(": ", 2)[0].replace("Card ", "")
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
                    wins = wins + 1

        orgCounterEl = amountCards[rowCounter-1][1] 
        
        for y in range(0, orgCounterEl):
            for x in range(rowCounter + 1, rowCounter +1  + wins):
                if x < numberOfCards:
                    amountCards[x-1][1] = amountCards[x-1][1]+1

        rowCounter = rowCounter + 1
   
    for i in amountCards:
        sumOfPoints = sumOfPoints + i[1]
    print(sumOfPoints)

#mainOne()
mainTwo()