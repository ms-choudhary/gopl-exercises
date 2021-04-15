# gopl-exercises
Go Programming Language Book Exercises


Solutions for following excersies:

##### Exercise 8.2
Implement a concurrent File Transfer Protocol (FTP) server. The
server should interpret commands from each client such as cd to change
directory, ls to list a directory, get to send the contents of a file, and close
to close the connection. You can use the standard ftp command as the client, or
write your own.

##### Exercise 8.5
Take an existing CPU-bound sequential program, such as the
Mandelbrot program of Section 3.3 or the 3-D surface computation of Section 3.2,
and execute its main loop in parallel using channels for communication. How much
faster does it run on a multiprocessor machine? What is the optimal number of
goroutines to use?

##### Exercise 8.7
Write a concurrent program that creates a local mirror of a web
site, fetching each reachable page and writing it to a directory on the local
disk. Only pages within the original domain (for instance, golang.org) should be
fetched. URLs within mirrored pages should be altered as needed so that they
refer to the mirrored page, not the original.
