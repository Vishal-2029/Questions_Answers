## Queue 

  - What is Queue?
     `Queue is a line collection of data item that follow FIFO principal. Where "Enter" and delere take place on opposite end. the Enter end is called "REAR" and delete end is called "Front".`

  - Enqueue                        - Dequeue
    ` '-----> For Enter of item `    ` '-----> For Delete of item `

  Types of Implementation 
        1. Array
        2. Linked list 

 ```bash
  Real Time example                            |                Programing Example 
                                               |
          ____                                 |         _____________________________________
         | ** | <--- Ticket Window             |        |    |  10  |  20  |  30  |  40  |    |
         |_  _|                                |         ______________________________________  
           *    <----Front (FIFO Principle)    |          0     1       2      3      4     5
           *                                   |
           *                                   |        Front = 1
           *    <----Rear                      |        Rear  = 4
                                               |        N     = 6 (Maximam Queue)
```
