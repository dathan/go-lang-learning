
```
Backtracking is an algorithm that finds candidate solutions and rejects a candidate on the basis of its feasibility and validity. Backtracking is useful in scenarios such as finding a value in an unordered table. It is faster than a brute force algorithm, which rejects a large number of solutions in an iteration. Constraint satisfaction problems such as parsing, rules engine, knapsack problems, and combinatorial optimization are solved using backtracking.

```

Another explination

```
Backtracking is an effective technique for solving algorithmic problems. In backtracking, we search depth-first for solutions, backtracking to the last valid path as soon as we hit a dead end.

Backtracking reduces the search space since we no longer have to follow down any paths we know are invalid. This is called pruning. We must be able to test partial solutions: for example, we can't find a global optimum using backtracking, since we have no idea if the solution we're currently on can lead to it or not. But we can, for example, solve Sudoku using backtracking. We can know immediately if our solution so far is invalid by testing if two of the same number appear in the same row, column, or square.

Let's go through several examples of problems that can be nicely solved with backtracking to drill this concept down.
```
