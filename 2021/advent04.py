import aocd

import numpy as np
from pprint import pprint

with open("data/04.txt", "r") as file:
    test_data = file.read()

prod_data = aocd.get_data(day=4, year=2021)


def init_matrix(x, y, value):
    return [[value] * x for _ in range(y)]


def get_board(data):
    board = init_matrix(x=5, y=5, value=(0, False))
    for y, row in enumerate(data.split("\n")):
        for x, col in enumerate(row.split()):
            board[y][x] = (col, False)
    return board


def get_boards(data):
    sections = data.split("\n\n")
    return list(map(lambda x: get_board(x), sections[1:]))


def play_number(board, number):
    for y, row in enumerate(board):
        for x, col in enumerate(row):
            if col[0] == number:
                board[y][x] = (number, True)
    return board


def check_board(board):
    b = np.array(board, object)
    for i in range(len(board)):
        row = b[i, :]
        col = b[:, i]

        row_check = all(mark for (_, mark) in row)
        col_check = all(mark for (_, mark) in col)

        if row_check or col_check:
            return True
    return False


def count_score(board, number):
    sum = 0
    for y, row in enumerate(board):
        for x in range(len(row)):
            if board[y][x][1] == False:
                sum += int(board[y][x][0])
    return sum * int(number)


def play_bingo(data):
    boards = get_boards(data)
    for number in data.split("\n")[0].split(","):
        for board in boards:
            play_number(board, number)
            if check_board(board):
                return count_score(board, number)


print("[TEST]:", play_bingo(test_data))
print("[PROD]:", play_bingo(prod_data))
