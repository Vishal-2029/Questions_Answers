﻿# UNIX_Signs "|" AND ">"

 ## "|" (Pipe) Sign 

     This Sign is use for redirect the Output one command as a input another command. this allowed for sequence of commands to be executed, where each command process the output of the previous one.
     
  Example
  
      - ls | grep file.txt - This command lists the files in the directory and pipes the output to grep, which searches for lines containing "file.txt".
      
  Use of Pipe 

    - command-to-command data flow 
    - allowing complex operations by combining multiple commands

 ## ">" (Redirection) Sign 

     This Sign is use for Output redirection,sending the output of a command to a file instead of displaying it on the terminal.
     
  Example
  
      - ls > log.txt - This command redirects the output of the ls command to a file named "log.txt", overwriting any existing content.
      
  Use of Pipe 

    - save command outputs to files for later reference or further processing.
 
