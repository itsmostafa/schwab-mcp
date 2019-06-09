'''
Bob the adventurer is one step away from solving the mystery of an ancient Mayan tomb.
He just approached the secret chamber where the secret Mayan scriptures are locked in a chest.

There are N ancient statues in the room. After long thought, Bob figured out that in order to open the treasure chest,
he needs to stand in the middle of the room and hit every statue with a laser ray at the same time.
Bob is highly experienced adventurer, so setting multiple laser rays at the same time is not a problem for him.
Moreover, every ray that he creates is perfectly straight and never changes direction at all.

The middle of the room, where Bob is standing, has coordinates (0,0). Every statue is located at some point with coordinates (x, y).
Each statue is made of pure glass, so that if any ray hits it, it does not stop, but goes through the statue and continues beyond in
the same, unchanged direction.

Bob wonders how he can hit every ancient statue in the room using the fewest rays possible.

Assume that the following declarations are given:

class Point2D(object):
    x = 0
    y = 0

Write a function

    def solution(A)

that, given an array of points A, representing the locations of the status,
returns the minimal number of rays that Bob must set in order to hit every statue in the room.

For example, given an array A

A[0].x = -1;  A[0].y = -2  (statue 0)
A[1].x = 1;   A[1].y = 2   (statue 1)
A[2].x = 2;   A[2].y = 4   (statue 2) 
A[3].x = -3;  A[3].y = 2   (statue 3)
A[4].x = 2;   A[4].y = -2  (statue 4)

your function should return 4.

As shown in the image, bob_the_adventurer.png, it is possible to create four rays in such a way that:

    - the first will hit statue 0;
    - the second will hit statues 1 and 2;
    - the third will hit statue 3;
    - the fourth will hit statue 4;

Write an efficient algorithm with the following assumptions:
    - N is an integer with range [1.. 100,000];
    - the coordinates of each point in array A
    are integers within the range
    [-1,000,0000,000.. 1,000,000,000]
    - Array A does not contain point (0,0)
'''