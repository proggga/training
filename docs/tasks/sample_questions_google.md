## Sample questions from Programming Pearls

Given a file containing at most ten million 7-digit integers with no duplicates. What is an efficient way to print these numbers in ascending order using just 1.5Mb RAM and reading the data just once? What are the consequences of only having 1Mb of RAM and no other storage? How would your answer change if duplicates were permitted?

Given a sequential file that contains at most four billion 32-bit integers in random order, find a 32-bit integer that isn’t in the file (and there must be at least one missing…why?). How would you solve this with unlimited main memory? How would you solve it if you could use several external files but only a few bytes of main memory?

Rotate a one-dimensional vector of n elements left by i positions. For instance, with n=8 and i=3, the vector abcdefg is rotated to defghabc. Simple code uses an n-element intermediate vector to do the job in n steps. Can you rotate the vector in time proportional to n using only a few dozen extra bytes of storage?

Given a dictionary of English words, find sets of anagrams. For instance, “pots”, “stop”, and “tops” are all anagrams of one another because each can be found by permuting the letters of the others.

Write functions for the following date problems: given two dates, compute the number of days between them; given a date, return it’s day of the week; given a month and year, produce a calendar of the month as an array of characters

Given a very long sequence (say, billions or trillions) of bytes, how would you efficiently count the total number of one bits? (i.e. how many bits are turned on in the entire sequence)

Although Quicksort uses only O(logn) stack space on the average, it can use linear space in the worst case. Explain why, then modify the program to use only logarithmic space in the worst case.

Write a program for finding the kth-smallest element in the array x[0…n-1] in O(n) expected time. Your algorithm may permute the elements of x.

Build the fastest possible complete function to generate a sorted array of random integers without duplicates. (You need feel constrained to use any prescribed interface for set representation)

Implement heap-based priority queues to run as quickly as possible; at what values of n are they faster than sequential structures?