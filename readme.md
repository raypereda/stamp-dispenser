Here's a classic problem in dynamic programming. 

Suppose we have stamps that have denominations specified in a set D. For example, D = {100, 25, 5, 1}. 
For a given postage P, find the minimum number of stamps for that cover than postage. 

The postage dispensing problem is equivalent to making change problem where you find the minimum number of coins for the given amount of money.

Here is an example of dispensing stamps. If the postage is 32, we cover that with four stamps: one 25 cent, one 5 cent, and two 1 cent stamps. 

Here is the recurrence equation to guide the dynamic programming solution. Let stamps(P) be the minimum number of stamps needed to cover the postage P.

stamps(P) = mininum of stamps(P - D<sub>i</sub>) + 1
            where D<sub>i</sub> is one of the stamp denominations.

With a proof by induction, we prove that current relation applies for all P.

Base case: P = 1. You can cover that with the smallest denomination stamp.

Inductive case P > 1: Let's prove this by contradiction. Suppose that the there is a better than in the recurrent equation claims.

better_stamps(P) < min stamp(P - D<sub>i</sub>) + 1

Every solution with more than a stamp can be broken into at least two stamps. 

So, better_stamps(P) is at least 2. So, 

2 < min stamp(P - D<sub>i</sub>) + 1

1 < min stamp(P - D<sub>i</sub>)

But that is not true because we know that some optimal solutions have exactly 1 stamp.

So, the inductive case is true by contradiction.
