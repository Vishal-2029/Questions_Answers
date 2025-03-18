## Architecture 

 ## What is Monolithic Architecture?
   
   `If We have develop any application, than we develop those application as asinle unit.That mans Thay are tightly integrated into a single codebase.`

  - Benifit of Monolithic Architecture.

    `Easier to Develop`,`Easier to Manage` etc...

  - Disadvantages

    `Redeployment`,`Sacling Limitations`,`Lots of dependence` etc...

 ## What is Microsevices Architecture?

   `Microservices are an architectural approach to developing software applications as a collection of small, independent services that communicate with each other over a network.`

  -  Benifit of Microsevices Architecture

   `Independent Deployment`,`Flexible Scaling`,`technology Flexibility` etc..

  - How to interact with each other?

    ```
     Synchronous Communication                 Asynchronous Communication
      |  |                                         |   |
      |  '----> HTTP/REST APIs                     |   '----> Message Brokers
      '----> gRPC                                  '----> Event-Driven Architecture
       etc...                                      etc...
    ```

    - Disadvantages

      `Complex To develop`,`management overhead`,`High Infra Costs` etc..

      
