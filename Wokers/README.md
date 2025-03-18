## Queue Worker pattern

   `The Queue Worker pattern is a design pattern used to handle background tasks asynchronously using a queue system.`

  - How to Workes
    ```
     client ---> HTTP API ---> Queue ---> Worker ---> DB
    ```

    - A client sent the request on the API.
    - Those API is push the request into a queue.
    - The queue is hold the request.
    - A worker retrieves a task, process it and marks it as complete.
    - when worker completes processing, it removes the task from the queue.

   ## For Example 

     - When a customer places an order, an "Order Placed" event is pushed into a queue.
     - The message queue (e.g., RabbitMQ, Kafka, AWS SQS) holds the order request.
     -A worker process (e.g., a background job) retrieves the order and processes it:
              -> Confirms the payment.
              -> Sends an email.
              -> Updates stock.
     -Once the worker completes processing, it removes the task from the queue.
