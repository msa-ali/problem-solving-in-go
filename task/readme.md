# Problem Statement

Consider the below scenario

```js
A --|
    |-- D --|
B --|       |-- E
            |
C ----------|
```

Each node is a async job, illustrated by setTimeout or any other async operation.

- A, B, and C can run at the same time.
- D needs to wait for A and B to be done.
- E needs to wait for C and D to be done.

Implement an interface to take care of this scenario.
