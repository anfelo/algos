# Runtime Complexity

Describes the performance of an algorithm

- How much more processing power/time is required to run your algorithm if we doble the inputs?
- Without going into details of an algorithm, just in a high level view.

## Constant Time: 1

- No matter how many elements we're working with, the algorithm/operation will take the same amount of time to execute the algorithm

## Logarithmic Time: log(n)

- You have this if doubling the number of elements you are iterating over doesn't double the amount of work. Always assume that **searching** operations are log(n).

## Linear Time: n

- Iterating throught all elements in a collection of data. If you see a for loop spanning from '0' to 'array.length', you probably have 'n', or linear runtime.

## Quasilinear Time: n\*log(n)

- You have this if doubling the number of elements you are iterating over doesn't double the amount of work. Always assume that any **sorting** operation is n\*log(n)

## Quadratic Time: n^2

- Every element in a collection has to be compared to every other element. 'The handshake problem'

## Exponential Time: 2^n

- If you add a "single" element to a collection, the processing power required doubles. Try to avoid algorithms with this complexity.

# Big 'O' Notation

Big O is the way we analyze how efficient algorithms are (or code in this case) without getting too mired in the details. We can model how much time any function is going to take given n inputs (think an array of length n), but in reality we're interested in the order of magnitude of the number and not necessarily of the exact figure.

Another way of express the runtime complexity:

- O(n) -> Linear
- O(1) -> Constant
- O(n^2) -> Quadratic

## Why use Big O

For starters it is asked in interviews. Secondly, is is useful for taking a high level view and deciding if an implementation is going to match the performance profile that you need.

## Identifying Runtime Complexity

| Problem | Complexity |
| -- | -- |
| Iterating with a simple for loop through a single collection? | Probably O |
| Iterating through half a collection? | Still O(n). There are no constants in runtime |
| Iterating through two _different_ collections with separate for loops? | O(n+m) |
| Two nested for loops iterating over the same collection? | O(n^2) |
| Two nested for loops iterating over different collections? | O(n\*m) |
| Sorting? | O(n\*log(n)) |
| Searching a sorted array? | O(log(n)) |


# Space Complexity is a thing too

How much more memory is required by doubling the problem set? Normally it only matters on low end devices where the disk storage or RAM are limited.

# Trade Offs

* There are no rules. "Always do blank". Everything has context. These are just tools and loose decision-making frameworks for you to use to make a contextually good choice.
* There are frequently multiple good choices and almost never a perfect, "right" choice.
* Remember how I said that Big O allows you to ignore the coefficients (the 3 in the 3x²)? Sometimes those actually end up being super important. Big O, again, is just a broad stroke. Sometimes the details are super important.
* Just as frequently, even the broad strokes are super unimportant. If you have a function that is called just once a day as a background job, it doesn't matter if it finishes in 300 milliseconds or 30 seconds (probably, again, context is important.) Don't spin your wheels on unimportant things.
* Readability and maintainability are the most important things about code. Code is communication. Clever, performant code is fun to write but hard to maintain later when you have to go figure out what the hell you actually wrote. We write code so that later humans can understand it and secondarily so computers can execute it. If it was just for computers we'd all still be writing assembly. Write your code like you were writing a letter to a future human (probably yourself) on how this works.
* Taking the above into account, err on the side of simple code. Simple code is easier to maintain because you can understand it easier and typically ends up having less bugs.
* Human time is almost always more valuable than computer time.
* Normally it's a good idea to not prematurely optimize code. As a general principle, I try to have a perf problem before I try to solve it. Premature optimization will cause you a lot of problems. Frequently you're not solving the right problem and you're left when harder-to-deal-with code.
* 99% of the time you want to use the built-in features to a language or an existing module to do these sorts of heavy lifting. Rarely are you going to write your own sort, you'll just call .sort(). Usually your implementation won't be faster because the built-ins can do tricks you can't (like run it in C/Rust) and they tend to have far less bugs because so many people use them.