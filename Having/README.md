## Having & Where Difference

 ## What is Having?
    
    - The Having is use for filter on the result of `GROUP BY` based on the specific condition.

 ## What is Where?
    
    - The where is use for a add condition on slected columns, where as the having is a place condition on group created by the `GROUP BY`. 

 ## Explaination of Code

    - In this case i have a add the cvs file in my sql databse using Golang. then In mysql terminal using the `HAVING` and `WHERE` clause. 

 ## Run the code :

       go run main.go

  In mysql Terminal: 
    
  - `HAVING` clause
    
    ```bash
      SELECT gender, Medium, COUNT(*) AS count
      FROM student_performance
      GROUP BY gender, Medium
      HAVING COUNT(*) >= 10 AND COUNT(*) < 50;
    ```

  - `WHERE` clause
    
    ```bash
      SELECT gender, Medium, COUNT(*) AS count
      FROM student_performance
      where Performance = 'vg'
      GROUP BY gender, Medium;
    ```

    
