# Cog
Cog is a project to shine a light on how systems work, and why performance is important.

## So what does it do?
Cog provides a virtual runner which is a simple "processor node" which gets access to simulated resources (memory, storage, networking between other nodes). You are given access to an assembly-like language and a higher-level language which gets "compiled" into the "assembly". Instructions in cog have two types: timed and interfaced. Timed instructions are everything that relate to your program running. Interfaced are instructions that interface back with the cog process and give allow you to time the execution (in "processor cycles") of blocks or entire programs. When a program finishes execution it reports extensive statistics on what it did, and what it spent time doing. Cog also provides an interface for threaded programs, allowing you to simulate concurrency.

## How does it do that?
Cog will implement a simple "processor" which will be able to do basic binary operations, which then gets expanded to your larger program with more wide representations of data. In order to keep results consistent against runs and across devices, processors will time themselves and transform their execution time to account for the performance of the device they run on.

## The Cog Model
- Cog
    - Maximum number of users*
    - Maximum concurrent processors*
    - User
        - Runner
            - Number of processors*
            - Total memory available*
        - Storage
            - Past trials
            - Runner storage* (Select option A or B)
                - Option A
                    - Number of object records*
                    - Maximum record size*
                - Option B
                    - Maximum size of storage*

\* = Configurable

## How do I run this?
Simply compile the program (this assumes you have at least Go 1.23.0 installed and working), and then provide it an environment file. Then run the binary output. <br/>

I recommend the following specs in order to run Cog efficiently with multiple users:
- 8 OS Threads at at least 3.000GHz
- 16 GB Memory
- 8 GB Storage