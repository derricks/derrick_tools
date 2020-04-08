from random import randint
from time import sleep

letters = ["a","ab","cd","e", "fg","hij","klmn"]
def generate_data_row():
    letter_index = randint(0,len(letters) - 1)
    data_point = randint(1,20)

    print(letters[letter_index] + " " + str(data_point), flush=True)

if __name__ == '__main__':
    while True:
        generate_data_row()
        sleep(1)
