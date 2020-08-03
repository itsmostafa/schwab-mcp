from random_words import RandomWords
import random


def letter_occurence(letter):
    indices = [i for i,x in enumerate(word) if x == letter]
    return indices


def player_choice(choice):
    board[pos] = letter
    
    if len(choice) == 1 and choice in word:
        print('yes')
        indices = letter_occurence(choice)
        if len(indices) > 1:
            for i in indices:
                board[i] = choice
        else:
            board[word.index(choice)] = choice
        print(board)
    elif choice in board:
        board[word.index(choice)] = choice
    elif choice == 'show1':
          print(word)
    elif choice == word:
          print('full word guessed without any hint') 
    else: 
        print('Incorrect')

def setup_board():
    board = []
    for every_letter in word:
      board.append('_')
    return board
# set up
rw = RandomWords()
word = rw.random_word()
pos = random.randint(0,len(word))
letter = word[pos]
board = setup_board()
# start
print(board)

print(f'The word is {len(word)} letters long, and the {pos+1}th letter is {letter}')

while '_' in board:
    player_choice(input('Choose a letter: '))
if '_' not in board:
    print(f'You got the word! It is {word}')
