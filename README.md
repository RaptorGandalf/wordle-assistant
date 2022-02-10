# wordle-assistant
Impress your friends with your ability to maybe solve the Wordle most of the time (probably). 

This was coded as quickly and dirtily as possible in a few hours. No guarantees of fame, riches or passage to Stovokor are provided with this software.

## Usage

The program runs a loop where you will

1. Enter new letters you know are not in the word.
2. Enter the letters you do know in a special format.
3. Recieve a best guess suggestion which attempts to find a remaining word with the highest number of unique characters to try and maximize information from that guess.

### Entering Known Letters

Known letters should be entered in sets of five.
If you don't know the letter in a given position place an asterisk instead.
If you know a letter is in the correct position enter it uppercase.
If you know the letter is in an incorrect position enter it lowercase.

For example, given the word you are trying to get to is humor, input of `o***R` would tell the program that you know an O exists in the word and that it is NOT in first position, and that you know an R is in the word in the last position. 
