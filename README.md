Gosmo
=====

Modular Go System Monitor



**Where:** 

Small non-critical networks. We hope to support most free unix like operating system. 

**Why:**

There were no solution for small scale highly customizable system monitoring. And we like to learn Go.

**What:** 

* Gosmo/server: 
    * Listens to selected port and receives information from client
* Gosmo/server/alert:
    * Executes script / sends mail when certain criteria is met, for example: CPU load > 4.0 over 5 minutes. 
* Gosmo/server/save: 
    * Saves received and selected data to file or db
* Gosmo/client: 
    * Gathers basic information from system, for example: cpu load, memory usage and HDD space and sends it to Gosmo/server
    * .conf file to manage polling interval and information gathered / sent to Gosmo/server
* Gosmo/client/module: 
    * Gathers information about the system, for example: queries per minute on mysqld or CPU temperature
    * Gosmo will come with few example modules 
    * User can create new modules using scripting language or binaries






This project is a living entity and will change rapidly and irrationally.
