# Day 11: Dumbo Octopus

You enter a large cavern full of rare bioluminescent [dumbo octopuses](https://www.youtube.com/watch?v=eih-VSaS2g0)! They seem to not like the Christmas lights on your submarine, so you turn them off for now.

There are 100 octopuses arranged neatly in a 10 by 10 grid. Each octopus slowly gains **energy** over time and **flashes** brightly for a moment when its energy is full. Although your lights are off, maybe you could navigate through the cave without disturbing the octopuses if you could predict when the flashes of light will happen.

Each octopus has an **energy level** - your submarine can remotely measure the energy level of each octopus (your puzzle input). For example:

```
 5  4  8  3  1  4  3  2  2  3
 2  7  4  5  8  5  4  7  1  1
 5  2  6  4  5  5  6  1  7  3
 6  1  4  1  3  3  6  1  4  6
 6  3  5  7  3  8  5  4  7  8
 4  1  6  7  5  2  4  6  4  5
 2  1  7  6  8  4  1  7  2  1
 6  8  8  2  8  8  1  1  3  4
 4  8  4  6  8  4  8  5  5  4
 5  2  8  3  7  5  1  5  2  6
```

The energy level of each octopus is a value between `0` and `9`. Here, the top-left octopus has an energy level of `5`, the bottom-right one has an energy level of `6`, and so on.

You can model the energy levels and flashes of light in **steps**. During a single step, the following occurs:

- First, the energy level of each octopus increases by `1`.
- Then, any octopus with an energy level greater than `9` **flashes**. This increases the energy level of all adjacent octopuses by `1`, including octopuses that are diagonally adjacent. If this causes an octopus to have an energy level greater than `9`, it **also flashes**. This process continues as long as new octopuses keep having their energy level increased beyond `9`. (An octopus can only flash **at most once per step**.)
- Finally, any octopus that flashed during this step has its energy level set to `0`, as it used all of its energy to flash.

Adjacent flashes can cause an octopus to flash on a step even if it begins that step with very little energy. Consider the middle octopus with `1` energy in this situation:

```
Before any steps:
 1  1  1  1  1
 1  9  9  9  1
 1  9  1  9  1
 1  9  9  9  1
 1  1  1  1  1

After step 1:
 3  4  5  4  3
 4 *0**0**0* 4
 5 *0**0**0* 5
 4 *0**0**0* 4
 3  4  5  4  3

After step 2:
 4  5  6  5  4
 5  1  1  1  5
 6  1  1  1  6
 5  1  1  1  5
 4  5  6  5  4
```

An octopus is **highlighted** when it flashed during the given step.

Here is how the larger example above progresses:

```
Before any steps:
 5  4  8  3  1  4  3  2  2  3
 2  7  4  5  8  5  4  7  1  1
 5  2  6  4  5  5  6  1  7  3
 6  1  4  1  3  3  6  1  4  6
 6  3  5  7  3  8  5  4  7  8
 4  1  6  7  5  2  4  6  4  5
 2  1  7  6  8  4  1  7  2  1
 6  8  8  2  8  8  1  1  3  4
 4  8  4  6  8  4  8  5  5  4
 5  2  8  3  7  5  1  5  2  6

After step 1:
 6  5  9  4  2  5  4  3  3  4
 3  8  5  6  9  6  5  8  2  2
 6  3  7  5  6  6  7  2  8  4
 7  2  5  2  4  4  7  2  5  7
 7  4  6  8  4  9  6  5  8  9
 5  2  7  8  6  3  5  7  5  6
 3  2  8  7  9  5  2  8  3  2
 7  9  9  3  9  9  2  2  4  5
 5  9  5  7  9  5  9  6  6  5
 6  3  9  4  8  6  2  6  3  7

After step 2:
 8  8 *0* 7  4  7  6  5  5  5
 5 *0* 8  9 *0* 8  7 *0* 5  4
 8  5  9  7  8  8  9  6 *0* 8
 8  4  8  5  7  6  9  6 *0**0*
 8  7 *0**0* 9 *0* 8  8 *0**0*
 6  6 *0**0**0* 8  8  9  8  9
 6  8 *0**0**0**0* 5  9  4  3
*0**0**0**0**0**0* 7  4  5  6
 9 *0**0**0**0**0**0* 8  7  6
 8  7 *0**0**0**0* 6  8  4  8

After step 3:
*0**0* 5 *0* 9 *0**0* 8  6  6
 8  5 *0**0* 8 *0**0* 5  7  5
 9  9 *0**0**0**0**0**0* 3  9
 9  7 *0**0**0**0**0**0* 4  1
 9  9  3  5 *0* 8 *0**0* 6  3
 7  7  1  2  3 *0**0**0**0**0*
 7  9  1  1  2  5 *0**0**0* 9
 2  2  1  1  1  3 *0**0**0**0*
*0* 4  2  1  1  2  5 *0**0**0*
*0**0* 2  1  1  1  9 *0**0**0*

After step 4:
 2  2  6  3 *0* 3  1  9  7  7
*0* 9  2  3 *0* 3  1  6  9  7
*0**0* 3  2  2  2  1  1  5 *0*
*0**0* 4  1  1  1  1  1  6  3
*0**0* 7  6  1  9  1  1  7  4
*0**0* 5  3  4  1  1  1  2  2
*0**0* 4  2  3  6  1  1  2 *0*
5  5  3  2  2  4  1  1  2  2
 1  5  3  2  2  4  7  2  1  1
 1  1  3  2  2  3 *0* 2  1  1

After step 5:
 4  4  8  4  1  4  4 *0**0**0*
 2 *0* 4  4  1  4  4 *0**0**0*
 2  2  5  3  3  3  3  4  9  3
 1  1  5  2  3  3  3  2  7  4
 1  1  8  7  3 *0* 3  2  8  5
 1  1  6  4  6  3  3  2  3  3
 1  1  5  3  4  7  2  2  3  1
 6  6  4  3  3  5  2  2  3  3
 2  6  4  3  3  5  8  3  2  2
 2  2  4  3  3  4  1  3  2  2

After step 6:
 5  5  9  5  2  5  5  1  1  1
 3  1  5  5  2  5  5  2  2  2
 3  3  6  4  4  4  4  6 *0* 5
 2  2  6  3  4  4  4  4  9  6
 2  2  9  8  4  1  4  3  9  6
 2  2  7  5  7  4  4  3  4  4
 2  2  6  4  5  8  3  3  4  2
 7  7  5  4  4  6  3  3  4  4
 3  7  5  4  4  6  9  4  3  3
 3  3  5  4  4  5  2  4  3  3

After step 7:
 6  7 *0* 7  3  6  6  2  2  2
 4  3  7  7  3  6  6  3  3  3
 4  4  7  5  5  5  5  8  2  7
 3  4  9  6  6  5  5  7 *0* 9
 3  5 *0**0* 6  2  5  6 *0* 9
 3  5 *0* 9  9  5  5  5  6  6
 3  4  8  6  6  9  4  4  5  3
 8  8  6  5  5  8  5  5  5  5
 4  8  6  5  5  8 *0* 6  4  4
 4  4  6  5  5  7  4  6  4  4

After step 8:
 7  8  1  8  4  7  7  3  3  3
 5  4  8  8  4  7  7  4  4  4
 5  6  9  7  6  6  6  9  4  9
 4  6 *0* 8  7  6  6  8  3 *0*
 4  7  3  4  9  4  6  7  3 *0*
 4  7  4 *0**0* 9  7  6  8  8
 6  9 *0**0**0**0* 7  5  6  4
*0**0**0**0**0**0* 9  6  6  6
 8 *0**0**0**0**0* 4  7  5  5
 6  8 *0**0**0**0* 7  7  5  5

After step 9:
 9 *0* 6 *0**0**0**0* 6  4  4
 7  8 *0**0**0**0**0* 9  7  6
 6  9 *0**0**0**0**0**0* 8 *0*
 5  8  4 *0**0**0**0**0* 8  2
 5  8  5  8 *0**0**0**0* 9  3
 6  9  6  2  4 *0**0**0**0**0*
 8 *0* 2  1  2  5 *0**0**0* 9
 2  2  2  1  1  3 *0**0**0* 9
 9  1  1  1  1  2  8 *0* 9  7
 7  9  1  1  1  1  9  9  7  6

After step 10
*0* 4  8  1  1  1  2  9  7  6
*0**0* 3  1  1  1  2 *0**0* 9
*0**0* 4  1  1  1  2  5 *0* 4
*0**0* 8  1  1  1  1  4 *0* 6
*0**0* 9  9  1  1  1  3 *0* 6
*0**0* 9  3  5  1  1  2  3  3
*0* 4  4  2  3  6  1  1  3 *0*
 5  5  3  2  2  5  2  3  5 *0*
*0* 5  3  2  2  5 *0* 6 *0**0*
*0**0* 3  2  2  4 *0**0**0**0*
```

After step 10, there have been a total of `204` flashes. Fast forwarding, here is the same configuration every 10 steps:

```
After step 20:
 3  9  3  6  5  5  6  4  5  2
 5  6  8  6  5  5  6  8 *0* 6
 4  4  9  6  5  5  5  6  9 *0*
 4  4  4  8  6  5  5  5  8 *0*
 4  4  5  6  8  6  5  5  7 *0*
 5  6  8 *0**0* 8  6  5  7  7
 7 *0**0**0**0**0* 9  8  9  6
*0**0**0**0**0**0**0* 3  4  4
 6 *0**0**0**0**0**0* 3  6  4
 4  6 *0**0**0**0* 9  5  4  3

After step 30:
*0* 6  4  3  3  3  4  1  1  8
 4  2  5  3  3  3  4  6  1  1
 3  3  7  4  3  3  3  4  5  8
 2  2  2  5  3  3  3  3  3  7
 2  2  2  9  3  3  3  3  3  8
 2  2  7  6  7  3  3  3  3  3
 2  7  5  4  5  7  4  5  6  5
 5  5  4  4  4  5  8  5  1  1
 9  4  4  4  4  4  7  1  1  1
 7  9  4  4  4  4  6  1  1  9

After step 40:
 6  2  1  1  1  1  1  9  8  1
*0* 4  2  1  1  1  1  1  1  9
*0**0* 4  2  1  1  1  1  1  5
*0**0**0* 3  1  1  1  1  1  5
*0**0**0* 3  1  1  1  1  1  6
*0**0* 6  5  6  1  1  1  1  1
*0* 5  3  2  3  5  1  1  1  1
 3  3  2  2  2  3  4  5  9  7
 2  2  2  2  2  2  2  9  7  6
 2  2  2  2  2  2  2  7  6  2

After step 50:
 9  6  5  5  5  5  6  4  4  7
 4  8  6  5  5  5  6  8 *0* 5
 4  4  8  6  5  5  5  6  9 *0*
 4  4  5  8  6  5  5  5  8 *0*
 4  5  7  4  8  6  5  5  7 *0*
 5  7 *0**0**0* 8  6  5  6  6
 6 *0**0**0**0**0* 9  8  8  7
 8 *0**0**0**0**0**0* 5  3  3
 6  8 *0**0**0**0**0* 6  3  3
 5  6  8 *0**0**0**0* 5  3  8

After step 60:
 2  5  3  3  3  3  4  2 *0**0*
 2  7  4  3  3  3  4  6  4 *0*
 2  2  6  4  3  3  3  4  5  8
 2  2  2  5  3  3  3  3  3  7
 2  2  2  5  3  3  3  3  3  8
 2  2  8  7  8  3  3  3  3  3
 3  8  5  4  5  7  3  4  5  5
 1  8  5  4  4  5  8  6  1  1
 1  1  7  5  4  4  7  1  1  1
 1  1  1  5  4  4  6  1  1  1

After step 70:
 8  2  1  1  1  1  1  1  6  4
*0* 4  2  1  1  1  1  1  6  6
*0**0* 4  2  1  1  1  1  1  4
*0**0**0* 4  2  1  1  1  1  5
*0**0**0**0* 2  1  1  1  1  6
*0**0* 6  5  6  1  1  1  1  1
*0* 5  3  2  3  5  1  1  1  1
 7  3  2  2  2  3  5  1  1  7
 5  7  2  2  2  2  3  4  7  5
 4  5  7  2  2  2  2  7  5  4

After step 80:
 1  7  5  5  5  5  5  6  9  7
 5  9  6  5  5  5  5  6 *0* 9
 4  4  8  6  5  5  5  6  8 *0*
 4  4  5  8  6  5  5  5  8 *0*
 4  5  7 *0* 8  6  5  5  7 *0*
 5  7 *0**0**0* 8  6  5  6  6
 7 *0**0**0**0**0* 8  6  6  6
*0**0**0**0**0**0**0* 9  9 *0*
*0**0**0**0**0**0**0* 8 *0**0*
*0**0**0**0**0**0**0**0**0**0*

After step 90:
 7  4  3  3  3  3  3  5  2  2
 2  6  4  3  3  3  3  5  2  2
 2  2  6  4  3  3  3  4  5  8
 2  2  2  6  4  3  3  3  3  7
 2  2  2  2  4  3  3  3  3  8
 2  2  8  7  8  3  3  3  3  3
 2  8  5  4  5  7  3  3  3  3
 4  8  5  4  4  5  8  3  3  3
 3  3  8  7  7  7  9  3  3  3
 3  3  3  3  3  3  3  3  3  3

After step 100:
*0* 3  9  7  6  6  6  8  6  6
*0* 7  4  9  7  6  6  9  1  8
*0**0* 5  3  9  7  6  9  3  3
*0**0**0* 4  2  9  7  8  2  2
*0**0**0* 4  2  2  9  8  9  2
*0**0* 5  3  2  2  2  8  7  7
*0* 5  3  2  2  2  2  9  6  6
 9  3  2  2  2  2  8  9  6  6
 7  9  2  2  2  8  6  8  6  6
 6  7  8  9  9  9  8  7  6  6
```

After 100 steps, there have been a total of **`1656`** flashes.

## Part 1

Given the starting energy levels of the dumbo octopuses in your cavern, simulate 100 steps. **How many total flashes are there after 100 steps?**

<details>
  <summary>The answer is...</summary>

  `1697`
</details>

## Part 2

It seems like the individual flashes aren't bright enough to navigate. However, you might have a better option: the flashes seem to be **synchronizing**!

In the example above, the first time all octopuses flash simultaneously is step **`195`**:

```
After step 193:
5877777777
8877777777
7777777777
7777777777
7777777777
7777777777
7777777777
7777777777
7777777777
7777777777

After step 194:
6988888888
9988888888
8888888888
8888888888
8888888888
8888888888
8888888888
8888888888
8888888888
8888888888

After step 195:
0000000000
0000000000
0000000000
0000000000
0000000000
0000000000
0000000000
0000000000
0000000000
0000000000
```

If you can calculate the exact moments when the octopuses will all flash simultaneously, you should be able to navigate through the cavern. **What is the first step during which all octopuses flash?**

<details>
  <summary>The answer is...</summary>

  `344`
</details>
