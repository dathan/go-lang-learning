from pprint import pprint

def n_queens(n, board=[]):
    if n == len(board):
        return 1

    count = 0
    for col in range(n):
        board.append(col)
        if is_valid(board):
            count += n_queens(n, board)
        board.pop()
    return count

def is_valid(board):
    current_queen_row, current_queen_col = len(board) - 1, board[-1]
    # Check if any queens can attack the last queen.
    # dump(enumerate(board[:-1]))
    for row, col in enumerate(board[:-1]): # next to last
        diff = abs(current_queen_col - col)
        if diff == 0 or diff == current_queen_row - row:
            return False
    return True

def dump(obj):
      for attr in dir(obj):
        print("obj.%s = %r" % (attr, getattr(obj, attr)))

for i in range(10):
    print(n_queens(i))