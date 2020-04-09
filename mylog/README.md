# Log Object

- The client can pass different parameters to the factory log function using a Configurator.
- There are 3 different type of log files that can be chosen by client.
 - #### File log
   - The sink of all log messages will be a log file.
   - The client can set log file size and number of log file rotations
 - #### Mem log
   - The sink of all log messages will be a chunk of memory.
 - #### Stdout log
   - The sink of all log messages will be standard output.

# Logger Object

- After creating a log a logger has to be obtained to write to the log
- A logger has a section (related to a section of code)
- Each section logger has its own log level
- The log level of all sections in the program can be obtained from config file in a string format
- A string format like "ALL:INFO, test1:ERR, test2:DBG" means all of the sections have default log level INFO, section test1 has log level ERR and section test2 has log level DBG.


